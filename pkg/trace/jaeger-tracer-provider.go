package trace

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// jaegerTracerProvider implements the TracerProvider interface using
// a Jaeger exporter.
type jaegerTracerProvider struct {
	providerSDK *tracesdk.TracerProvider
	tracerName  string
}

func newJaegerTracerProvider(config TracerProviderConfig) (TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.CollectorURL)))
	if err != nil {
		return nil, err
	}
	providerSDK := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.TracerService),
			attribute.String("environment", config.TracerEnvironment),
			attribute.Int64("ID", config.TracerID),
		)),
	)
	tp := jaegerTracerProvider{
		providerSDK: providerSDK,
		tracerName:  config.TracerName,
	}
	return &tp, nil
}

func (tp *jaegerTracerProvider) Tracer(name string) Tracer {
	return NewOtelTracer(tp.providerSDK.Tracer(tp.tracerName))
}

func (tp *jaegerTracerProvider) Shutdown(ctx context.Context) error {
	// Do not make the application hang when it is shutdown.
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	return tp.providerSDK.Shutdown(ctx)
}

func (tp *jaegerTracerProvider) SetGlobalTracerProvider() {
	otel.SetTracerProvider(tp.providerSDK)
}
