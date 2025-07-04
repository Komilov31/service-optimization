package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Producer struct {
	writer *kafka.Writer
	logger *zap.Logger
}

func NewProducer(logger *zap.Logger) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "orders",
	})

	return &Producer{writer: writer, logger: logger}
}

func (p *Producer) Produce() {
	defer p.writer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		p.logger.Info(fmt.Sprintf("recieved shutting signal %v. Shuting down producer", sig),
			zap.Time("time", time.Now()),
		)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			p.logger.Info("producer stopped")
			return
		default:
			order := GenerateNewOrder()
			msg, err := json.Marshal(&order)
			if err != nil {
				p.logger.Error(fmt.Sprintf("could not marshal order: %v\n", err))
				os.Exit(1)
			}

			err = p.writer.WriteMessages(ctx, kafka.Message{
				Value: msg,
			})
			if err != nil {
				p.logger.Error(fmt.Sprintf("could not send message to kafka: %v\n", err))
			}
			p.logger.Info(
				"new order was sent",
				zap.Time("time", time.Now()),
			)
		}

		time.Sleep(time.Second * 2)
	}
}
