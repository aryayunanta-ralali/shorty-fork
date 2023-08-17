// Package main
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/pkg/kafka"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type MessageTX struct {
	Name        string `json:"name"`
	ReferenceID string `json:"reference"`
}


func main() {
	cfg, e := appctx.NewConfig()

	if e != nil {
		logger.Fatal(fmt.Sprintf("config error : %s", e.Error()))
	}

	kp := kafka.NewProducer(&kafka.Config{
		Producer: kafka.ProducerConfig{
			TimeoutSecond:     cfg.Kafka.Producer.TimeoutSecond,
			RequireACK:        cfg.Kafka.Producer.RequireACK,
			IdemPotent:        cfg.Kafka.Producer.IdemPotent,
			PartitionStrategy: cfg.Kafka.Producer.PartitionStrategy,
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

	ctx := context.Background()

	msg := MessageTX{
		ReferenceID: "test-123",
		Name:        "coba-coba",
	}

	e = kp.Publish(ctx, &kafka.MessageContext{
		Value: util.DumpToString(util.DumpToString(msg)),
		Topic: "dgs-test", // topic name
		Verbose: true,
	})

	if e != nil {
		logger.Error(logger.SetMessageFormat("publish message error: ", e))
	}
}
