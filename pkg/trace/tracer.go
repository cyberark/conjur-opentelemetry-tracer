package trace

import (
	"context"

	traceotel "go.opentelemetry.io/otel/trace"
)

// Tracer is responsible for creating trace Spans.
type Tracer interface {
	// Start creates a span and a context.Context containing the newly-created
	// span.
	//
	// If the context.Context provided in `ctx` contains a Span then the
	// newly-created Span will be a child of that span, otherwise it will be a
	// root span.
	Start(ctx context.Context, spanName string) (context.Context, Span)
}

// otelTracer implements the Tracer interface based on an OpenTelemetry
// Tracer.
type otelTracer struct {
	tracerOtel traceotel.Tracer
}

func NewOtelTracer(tracerOtel traceotel.Tracer) otelTracer {
	return otelTracer{tracerOtel: tracerOtel}
}

func (t otelTracer) Start(ctx context.Context, spanName string) (context.Context, Span) {
	ctx, spanOtel := t.tracerOtel.Start(ctx, spanName)
	return ctx, newOtelSpan(spanOtel)
}

// Create returns a Context, Tracer and cleanup function based on the provided
// TracerProviderType and TracerProviderConfig.
func Create(
	providerType TracerProviderType,
	config TracerProviderConfig,
) (context.Context, Tracer, func(context.Context), error) {
	return create(
		providerType,
		config,
		NewTracerProvider,
	)
}

func create(
	providerType TracerProviderType,
	config TracerProviderConfig,
	factory func(TracerProviderType, bool, TracerProviderConfig) (TracerProvider, error),
) (context.Context, Tracer, func(context.Context), error) {
	tp, err := factory(
		providerType,
		SetGlobalProvider,
		config,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	tracer := tp.Tracer(config.TracerName)
	ctx, span := tracer.Start(ctx, "main")

	cleanupFunc := func(ctx context.Context) {
		span.End()
		tp.Shutdown(ctx)
		cancel()
	}

	return ctx, tracer, cleanupFunc, err
}
