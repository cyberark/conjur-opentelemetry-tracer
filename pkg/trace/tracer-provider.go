package trace

import (
	"context"
	"fmt"
	"io"
	"os"

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
	TracerName string
	// Service to be traced
	TracerService string
	// Execution environment such as "production" or "development"
	TracerEnvironment string
	// Unique ID of the tracer
	TracerID int64
	// URL of the collector when using Jaeger
	CollectorURL string
	// Writer to use for the console tracer
	ConsoleWriter io.Writer
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

// TypeFromAnnotationsOrEnv return a TracerProviderType and optional Jaeger
// collector URL based on a map of annotation key-value pairs and environment
// variable configuration. Annotations are prioritized over envvars.
func TypeFromAnnotationsOrEnv(annots map[string]string) (TracerProviderType, string) {
	return typeFromAnnotationsOrCustomEnv(annots, os.Getenv)
}

func typeFromAnnotationsOrCustomEnv(annots map[string]string, getEnv func(string) string) (TracerProviderType, string) {
	providerType, jaegerUrl := TypeFromAnnotations(annots)
	if providerType == NoopProviderType {
		providerType, jaegerUrl = typeFromCustomEnv(getEnv)
	}
	return providerType, jaegerUrl
}

// TypeFromEnv returns a TracerProviderType and optional Jaeger collector URL
// based on environment variable configuration.
func TypeFromEnv() (TracerProviderType, string) {
	return typeFromCustomEnv(os.Getenv)
}

func typeFromCustomEnv(getEnv func(string) string) (TracerProviderType, string) {
	jaegerUrl := getEnv("JAEGER_COLLECTOR_URL")
	if jaegerUrl != "" {
		return JaegerProviderType, jaegerUrl
	}
	if getEnv("LOG_TRACES") == "true" {
		return ConsoleProviderType, ""
	}
	return NoopProviderType, ""
}

// TypeFromAnnotations returns a TracerProviderType and optional Jaeger
// collector URL based on a map of annotations key-value pairs.
func TypeFromAnnotations(annots map[string]string) (TracerProviderType, string) {
	jaegerUrl := annots["conjur.org/jaeger-collector-url"]
	if jaegerUrl != "" {
		return JaegerProviderType, jaegerUrl
	}
	if annots["conjur.org/log-traces"] == "true" {
		return ConsoleProviderType, ""
	}
	return NoopProviderType, ""
}
