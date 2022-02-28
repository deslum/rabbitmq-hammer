package main

import (
	"flag"
	"log"

	"rabbitmq-hammer/config"
	"rabbitmq-hammer/logs"
	"rabbitmq-hammer/process"
	"rabbitmq-hammer/rmq"
	"rabbitmq-hammer/statistic"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(new(logs.LogWriter))

	// Read RMQHammer configuration file
	cfg, err := config.NewRMQHammerConfig().GetConfig()
	if err != nil {
		log.Println(err)
		return
	}

	var reader, writer process.Processor

	switch {
	// Consumer in config not empty
	case cfg.Consumer != nil:
		consumer, err := rmq.NewConsumer(cfg)
		if err != nil {
			log.Println(err)
			return
		}

		reader = process.NewReader(consumer, cfg)
		reader.Start()
		fallthrough
	// Publisher in config not empty
	case cfg.Publisher != nil:
		// Connect publisher to RabbitMQ
		publisher, err := rmq.NewPublisher(cfg.URI)
		if err != nil {
			log.Println(err)
			return
		}

		writer = process.NewWriter(publisher, cfg)
		writer.Start()
	}

	writerStat := statistic.NewStatistic(reader, writer)
	writerStat.Show()
}
