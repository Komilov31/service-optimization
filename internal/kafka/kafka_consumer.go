package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Komilov31/l0/internal/model"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Consumer struct {
	reader  *kafka.Reader
	logger  *zap.Logger
	service model.OrderService
}

func NewConsumer(logger *zap.Logger, service model.OrderService) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"kafka:9092"},
		Topic:          "orders",
		MaxBytes:       10e6,
		CommitInterval: 0,
		GroupID:        "first",
	})

	return &Consumer{
		reader:  reader,
		logger:  logger,
		service: service,
	}
}

func (c *Consumer) Consume() {
	defer c.reader.Close()

	c.logger.Info(
		"consumer started consuming messages from kafka",
		zap.Time("time", time.Now()),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		c.logger.Info(fmt.Sprintf("recieved shutting signal %v. Shuting down", sig),
			zap.Time("time", time.Now()),
		)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			c.logger.Info("consumer stopped")
			return
		default:
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				c.logger.Error(fmt.Sprintf("could not read message from kafka: %v", err))
				if err := c.reader.CommitMessages(ctx, msg); err != nil {
					c.logger.Error(
						"failed to commit offset",
						zap.String("errorMessage", err.Error()),
					)
				}
				continue
			}

			var order model.Order
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				c.logger.Error(
					"invalid message from kafka",
					zap.Time("time", time.Now()),
				)
				if err := c.reader.CommitMessages(ctx, msg); err != nil {
					c.logger.Error(
						"failed to commit offset",
						zap.String("errorMessage", err.Error()),
					)
				}
				continue
			}

			err = c.service.CreateOrder(ctx, order)
			if err != nil {
				c.logger.Error("could not save order from kafka to db")
				if err := c.reader.CommitMessages(ctx, msg); err != nil {
					c.logger.Error(
						"failed to commit offset",
						zap.String("errorMessage", err.Error()),
					)
				}
				continue
			}

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				c.logger.Error(
					"failed to commit offset",
					zap.String("errorMessage", err.Error()),
				)
				continue
			}

		}
	}
}
