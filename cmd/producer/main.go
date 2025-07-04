package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("could not initialize logger")
	}
	defer logger.Sync()

	producer := NewProducer(logger)

	producer.Produce()
}
