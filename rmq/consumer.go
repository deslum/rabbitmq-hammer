package rmq

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"

	"rabbitmq-hammer/config"
	"rabbitmq-hammer/consts"
)

// Consumer is RabbitMQ consumer.
type Consumer struct {
	ConsumeChan      chan amqp.Delivery
	connection       *amqp.Connection
	channel          *amqp.Channel
	deliveryChannels []<-chan amqp.Delivery
	countQueues      int
	wg               *sync.WaitGroup
}

// NewConsumer initializes connection for RabbitMQ consumer.
func NewConsumer(cfg *config.RMQHammerConfig) (*Consumer, error) {
	deliveryBufferChannel := make(chan amqp.Delivery, consts.DeliveryBufferChannel)

	connection, err := amqp.Dial(cfg.URI)
	if err != nil {
		errString := "amqp.Dial(%v) connect error: %v"
		return nil, fmt.Errorf(errString, cfg.URI, err)
	}

	channel, err := connection.Channel()
	if err != nil {
		errString := "connection.Channel() error: %v"
		return nil, fmt.Errorf(errString, err)
	}

	err = channel.Qos(cfg.Consumer.PrefetchCount, 0, false)
	if err != nil {
		errString := "channel.Qos() setting QOS: %v"
		return nil, fmt.Errorf(errString, err)
	}

	var deliveryChannels []<-chan amqp.Delivery
	for _, queue := range cfg.Consumer.Queues {
		deliveryChannel, err := channel.Consume(
			queue,
			consts.ConsumerName,
			cfg.Consumer.AutoAck,
			cfg.Consumer.Exclusive,
			cfg.Consumer.NoLocal,
			cfg.Consumer.NoWait,
			nil,
		)
		if err != nil {
			errString := "channel.Consume error: %v"
			return nil, fmt.Errorf(errString, err)
		}

		deliveryChannels = append(deliveryChannels, deliveryChannel)

	}

	return &Consumer{
		ConsumeChan:      deliveryBufferChannel,
		deliveryChannels: deliveryChannels,
		connection:       connection,
		channel:          channel,
		countQueues:      len(cfg.Consumer.Queues),
		wg:               &sync.WaitGroup{},
	}, nil
}

// StartConsume function start consume messages from rabbitmq queues
func (o *Consumer) StartConsume() {
	// TODO Testing this is shit
	for numQueue := 0; numQueue < o.countQueues; numQueue++ {
		o.wg.Add(1)
		go o.consume(numQueue)
	}

	o.wg.Wait()
}

func (o *Consumer) consume(numQueue int) {
	defer o.wg.Done()
	for {
		msgChan, ok := <-o.deliveryChannels[numQueue]
		if !ok {
			return
		}

		o.ConsumeChan <- msgChan
	}
}
