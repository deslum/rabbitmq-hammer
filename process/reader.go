package process

import (
	"log"
	"sync"
	"time"

	"rabbitmq-hammer/config"
	"rabbitmq-hammer/rmq"
)

type Reader struct {
	cfg      *config.RMQHammerConfig
	consumer *rmq.Consumer

	statistic []float64
	wg        *sync.WaitGroup
}

func NewReader(consumer *rmq.Consumer, cfg *config.RMQHammerConfig) *Reader {
	return &Reader{
		consumer: consumer,
		cfg:      cfg,
		wg:       &sync.WaitGroup{},
	}
}

func (o *Reader) Start() {
	go o.consumer.StartConsume()

	for i := 0; i < o.cfg.Consumer.ConsumersCount; i++ {
		o.wg.Add(1)
		go o.process()
	}

}

func (o *Reader) process() {
	defer o.wg.Done()
	for {
		before := time.Now().UTC()
		message, ok := <-o.consumer.ConsumeChan
		if !ok {
			break
		}

		_, err := o.cfg.CompressType.Encode(message.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		o.statistic = append(o.statistic, float64(time.Now().UTC().Sub(before).Microseconds())/10)

		time.Sleep(time.Duration(o.cfg.Consumer.LatencyInMicroseconds) * time.Microsecond)
		err = message.Ack(false)
		if err != nil {
			log.Println(err)
		}
	}
}

func (o *Reader) GetStatistic() []float64 {
	return o.statistic
}
