package rmq

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/streadway/amqp"
	"rabbitmq-hammer/consts"
)

// Publisher struct implemented RabbitMQ publisher
type Publisher struct {
	channel    *amqp.Channel
	connection *amqp.Connection
	blockChan  chan amqp.Blocking
	sync.Mutex
}

func NewPublisher(uri string) (*Publisher, error) {
	connection, err := amqp.Dial(uri)
	if err != nil {
		errString := "amqp.Dial(%s) error: %v"
		errMessage := fmt.Errorf(errString, uri, err)
		return nil, errMessage
	}

	channel, err := connection.Channel()
	if err != nil {
		errString := "connection.Channel(%s) error: %v"
		errMessage := fmt.Errorf(errString, uri, err)
		return nil, errMessage
	}

	return &Publisher{
		connection: connection,
		blockChan:  connection.NotifyBlocked(make(chan amqp.Blocking)),
		channel:    channel,
	}, nil
}

// Publish message to RabbitMQ exchange
func (o *Publisher) SendMessage(exchangeName, key string, msg []byte, ttl uint64) error {
	message := amqp.Publishing{
		Expiration:      strconv.FormatInt(int64(ttl), 10),
		Headers:         amqp.Table{},
		ContentType:     consts.JSONContentType,
		ContentEncoding: consts.EmptyContentEncoding,
		Body:            msg,
		DeliveryMode:    amqp.Persistent,
	}

	select {
	case <-o.blockChan:
		log.Println("Error =>> flow")
		return fmt.Errorf("Vse Huevo!!!")
	default:
		err := o.channel.Publish(
			exchangeName,
			key,
			false,
			false,
			message,
		)
		if err != nil {
			errString := "channel.Publish(%v, %v, %v), error: %v"
			return fmt.Errorf(errString, exchangeName, key, string(msg), err)
		}
	}

	return nil
}

// Close function close publisher connection
func (o *Publisher) Close() {
	if err := o.channel.Close(); err != nil {
		log.Println(err)
	}
	if err := o.connection.Close(); err != nil {
		log.Println(err)
	}
}
