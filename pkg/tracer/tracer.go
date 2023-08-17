// Package tracer
// @author Daud Valentino
package tracer

import (
	"context"
	"fmt"

	"github.com/aryayunanta-ralali/shorty/pkg/util"

	"github.com/opentracing/opentracing-go"
	otext "github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

type SpanOption struct {
	SpanType     string                  `json:"span_type"`
	ServiceName  string                  `json:"service_name"`
	ResourceName string                  `json:"resource_name"`
	OptionsTag   *map[string]interface{} `json:"options_tag"`
}

type traceConfig struct {
	serviceName string
}

var cfg traceConfig

// New sets the given apm service name.
func New(name string) *traceConfig {
	cfg = traceConfig{serviceName: name}
	return &cfg
}

// SpanStart starts a new query span from ctx, then returns a new context with the new span.
func SpanStart(ctx context.Context, eventName string) context.Context {
	_, ctx = opentracing.StartSpanFromContext(ctx, eventName)
	return ctx
}

// SpanFinish finishes the span associated with ctx.
func SpanFinish(ctx context.Context) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span.Finish()
	}
}

// SpanError adds an error to the span associated with ctx.
func SpanError(ctx context.Context, err error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		otext.Error.Set(span, true)
		span.LogFields(otlog.String("event", "error"), otlog.Error(err))
	}
}

// SpanStartRepositories starts and return value is a context.Context object built around
// the returned Span with span type function and service name is repositories.
//
// Example usage:
//
//    AwesomeFunction(ctx context.Context, ...) {
// 		 	ctx := tracer.SpanStartRepositories(ctx, "AwesomeFunction")
// 			defer tracer.SpanFinish(ctx)
//        	...
//    }
func SpanStartRepositories(ctx context.Context, eventName string) context.Context {
	spanOption := SpanOption{
		SpanType:     "function",
		ServiceName:  "repositories",
		ResourceName: fmt.Sprintf("repositories.%s", eventName),
	}
	opts := WithStartSpanOptions(spanOption)

	_, ctx = opentracing.StartSpanFromContext(ctx, eventName, opts...)
	return ctx
}

// WithStartSpanOptions set a options span (type, service name and resource name)
func WithStartSpanOptions(spanOption SpanOption) []opentracing.StartSpanOption {
	apmName := cfg.serviceName

	var spanopts []opentracing.StartSpanOption

	opts := append([]opentracing.StartSpanOption{
		opentracing.Tag{Key: "span.type", Value: spanOption.SpanType},
		opentracing.Tag{Key: "service.name", Value: fmt.Sprintf("%s.%s", util.ToString(apmName), spanOption.ServiceName)},
		opentracing.Tag{Key: "resource.name", Value: spanOption.ResourceName},
	}, spanopts...)

	return opts
}

// SpanStartUseCase starts and return value is a context.Context object built around
// the returned Span with span type function and service name is usecase.
//
// Example usage:
//
//    AwesomeFunction(ctx context.Context, ...) {
// 		 	ctx := tracer.SpanStartUseCase(ctx, "AwesomeFunction")
// 			defer tracer.SpanFinish(ctx)
//        	...
//    }
func SpanStartUseCase(ctx context.Context, eventName string) context.Context {
	spanOption := SpanOption{
		SpanType:     "function",
		ServiceName:  "usecase",
		ResourceName: fmt.Sprintf("usecase.%s", eventName),
	}
	opts := WithStartSpanOptions(spanOption)

	_, ctx = opentracing.StartSpanFromContext(ctx, eventName, opts...)
	return ctx
}