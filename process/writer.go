package process

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"rabbitmq-hammer/config"
	"rabbitmq-hammer/rmq"
)

type Writer struct {
	cfg       *config.RMQHammerConfig
	publisher *rmq.Publisher
	wg        *sync.WaitGroup

	timeNsChan chan float64

	msgsCount              uint64
	messagesProcessLimiter uint64
	statistic              []float64
}

// NewWriter fucntion init writer on RabbitMQ
func NewWriter(publisher *rmq.Publisher, cfg *config.RMQHammerConfig) *Writer {
	return &Writer{
		msgsCount: 0,
		cfg:       cfg,
		publisher: publisher,
		wg:        &sync.WaitGroup{},

		messagesProcessLimiter: cfg.Publisher.MessagesCount / 10,
		timeNsChan:             make(chan float64, 10),
		statistic:              make([]float64, 0),
	}
}

func (o *Writer) Start() {
	go o.collectStat()
	for hammer := 0; hammer < o.cfg.Publisher.HammersCount; hammer++ {
		o.wg.Add(1)
		go o.process()
	}

	o.wg.Wait()
}

func (o *Writer) process() {
	defer o.wg.Done()
	for {
		before := time.Now().UTC()

		message, err := o.cfg.CompressType.Encode(o.cfg.Publisher.Message)
		if err != nil {
			log.Printf("compressType.Encode(%v) error %v", o.cfg.Publisher.Message, err)
			break
		}

		err = o.publisher.SendMessage(
			o.cfg.Publisher.Exchange,
			o.cfg.Publisher.RoutingKey,
			message,
			o.cfg.Publisher.TTL,
		)
		if err != nil {
			log.Fatal(err)
		}

		timeNs := float64(time.Now().UTC().Sub(before).Microseconds())

		o.timeNsChan <- timeNs

		atomic.AddUint64(&o.msgsCount, 1)

		msgsCount := atomic.LoadUint64(&o.msgsCount)
		if msgsCount > o.cfg.Publisher.MessagesCount {
			break
		}

		if msgsCount%(o.messagesProcessLimiter) == 0 {
			log.Printf("%v messages processed\n", msgsCount)
		}
	}
}

func (o *Writer) collectStat() {
	for {
		stat, ok := <-o.timeNsChan
		if !ok {
			break
		}

		o.statistic = append(o.statistic, stat/float64(o.cfg.Publisher.HammersCount))
	}
}

func (o *Writer) GetStatistic() []float64 {
	return o.statistic
}
