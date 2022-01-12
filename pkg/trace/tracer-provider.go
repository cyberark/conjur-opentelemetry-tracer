package trace

import (
	"context"
	"fmt"
	"io"

	"go.opentelemetry.io/otel"
)

// TracerProviderType represents a type of TracerProvider
type TracerProviderType int64

// Valid values for TracerProviderType
const (
	NoopProviderType = iota
	ConsoleProviderType
	JaegerProviderType
)

// Boolean flags for indicating whether a TracerProvider should be set as
// the global TracerProvider
const (
	SetGlobalProvider     = true
	DontSetGlobalProvider = false
)

type TracerProviderConfig struct {
	// Name of the tracer
	tracerName string
	// Service to be traced
	tracerService string
	// Execution environment such as "production" or "development"
	tracerEnvironment string
	// Unique ID of the tracer
	tracerID int64
	// URL of the collector when using Jaeger
	collectorURL string
	// Writer to use for the console tracer
	consoleWriter io.Writer
}

// TracerProvider provides access to Tracers, which in turn allow for creation
// of trace Spans.
type TracerProvider interface {
	// Tracer creates an implementation of the Tracer interface.
	Tracer(tracerName string) Tracer
	// Shutdown and flush telemetry.
	Shutdown(ctx context.Context) error
	SetGlobalTracerProvider()
}

// NewTracerProvider creates a TracerProvider of a given type, and
// optionally sets the new TracerProvider as the global TracerProvider.
func NewTracerProvider(
	providerType TracerProviderType,
	setGlobalProvider bool,
	config TracerProviderConfig) (TracerProvider, error) {

	var tp TracerProvider
	var err error

	switch providerType {
	case NoopProviderType:
		tp = newNoopTracerProvider()
	case ConsoleProviderType:
		tp, err = newConsoleTracerProvider(config)
	case JaegerProviderType:
		tp, err = newJaegerTracerProvider(config)
	default:
		err = fmt.Errorf("invalid TracerProviderType '%d' in call to NewTracerProvider",
			providerType)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if setGlobalProvider {
		tp.SetGlobalTracerProvider()
	}

	return tp, nil
}

// GlobalTracer returns a Tracer using the registered global trace provider.
func GlobalTracer(tracerName string) Tracer {
	tracer := otel.Tracer(tracerName)
	return NewOtelTracer(tracer)
}
