// Package handler
package handler

import (
	"context"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	uContract "github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/awssqs"
)

// SQSConsumerHandler sqs consumer message processor handler
func SQSConsumerHandler(msgHandler uContract.MessageProcessor) awssqs.MessageProcessorFunc {
	return func(decoder *awssqs.MessageDecoder) error {
		return msgHandler.Serve(context.Background(), &appctx.ConsumerData{
			Body:        []byte(*decoder.Body),
			Key:         []byte(*decoder.MessageId),
			ServiceType: consts.ServiceTypeConsumer,
		})
	}
}
