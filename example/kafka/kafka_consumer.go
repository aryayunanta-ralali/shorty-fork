// Package example
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/pkg/kafka"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type MessageReceive struct {
	Name        string `json:"name"`
	ReferenceID string `json:"reference"`
}

func main() {

	cfg, e := appctx.NewConfig()

	if e != nil {
		logger.Fatal(fmt.Sprintf("config error : %s", e.Error()))
	}

	consumer := kafka.NewConsumerGroup(&kafka.Config{
		Consumer: kafka.ConsumerConfig{
			SessionTimeoutSecond: cfg.Kafka.Consumer.SessionTimeoutSecond,
			HeartbeatInterval:    cfg.Kafka.Consumer.HeartbeatIntervalMS,
			RebalanceStrategy:    cfg.Kafka.Consumer.RebalanceStrategy,
			OffsetInitial:        cfg.Kafka.Consumer.OffsetInitial,
		},
		Version:  cfg.Kafka.Version,
		Brokers:  strings.Split(cfg.Kafka.Brokers, ","),
		ClientID: cfg.Kafka.ClientID,
		SASL: kafka.SASL{
			Enable:    cfg.Kafka.SASL.Enable,
			User:      cfg.Kafka.SASL.User,
			Password:  cfg.Kafka.SASL.Password,
			Mechanism: cfg.Kafka.SASL.Mechanism,
			Version:   cfg.Kafka.SASL.Version,
			Handshake: cfg.Kafka.SASL.Handshake,
		},
		TLS: kafka.TLS{
			Enable: true,
		},
		ChannelBufferSize: cfg.Kafka.ChannelBufferSize,
	})

	consumer.Subscribe(&kafka.ConsumerContext{
		Topics:  []string{"dgs-test"},
		GroupID: "contoh",
		Context: context.Background(),
		Handler: MessageHandler,
	})

}

func MessageHandler(msg *kafka.MessageDecoder) {
	logger.Info(fmt.Sprintf("message body = %q, topic = %s, partition = %d, offset = %d", msg.Body, msg.Topic, msg.Partition, msg.Offset))

	// TODO : create process logic here
	data := MessageReceive{}
	json.Unmarshal(msg.Body, &data)

	util.DebugPrint(data)

	msg.Commit(msg)
}
