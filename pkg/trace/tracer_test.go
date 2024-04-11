package trace

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
)

func TestTracer(t *testing.T) {
	testCases := []struct {
		name         string
		providerType TracerProviderType
		collectorUrl string
		assertFunc   func(t *testing.T, tracer Tracer, tp TracerProvider, output *bytes.Buffer)
	}{
		{
			name:         "ConsoleProvider",
			providerType: ConsoleProviderType,
			assertFunc: func(t *testing.T, tracer Tracer, tp TracerProvider, output *bytes.Buffer) {
				// Check provider type
				assert.IsType(t, &consoleTracerProvider{}, tp)

				// Check output
				str := output.String()
				assert.Contains(t, str, "ConsoleProvider_normal")
				assert.Contains(t, str, "ConsoleProvider_error")
				assert.Contains(t, str, "some fake error")
				assert.Contains(t, str, "testAttr")
				assert.Contains(t, str, "testValue")
			},
		},
		{
			name:         "JaegerProvider",
			providerType: JaegerProviderType,
			collectorUrl: "",
			assertFunc: func(t *testing.T, tracer Tracer, tp TracerProvider, output *bytes.Buffer) {
				// Check provider type
				assert.IsType(t, &jaegerTracerProvider{}, tp)

				// Output should not be in stdout
				assert.NotContains(t, output.String(), "TRACING OUTPUT")
				assert.NotContains(t, output.String(), "JaegerProvider_normal")

				// TODO: Check Jaeger output (mock server similar to conjur authn tests?)
			},
		},
		{
			name:         "NoopProvider",
			providerType: NoopProviderType,
			assertFunc: func(t *testing.T, tracer Tracer, tp TracerProvider, output *bytes.Buffer) {
				// Check provider type
				assert.IsType(t, &noopTracerProvider{}, tp)

				// Output should not be in stdout
				assert.NotContains(t, output.String(), "TRACING OUTPUT")
				assert.NotContains(t, output.String(), "NoopProvider_normal")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create an in memory writer to capture any console output
			output := &bytes.Buffer{}
			ctx := context.Background()

			tp, err := NewTracerProvider(tc.providerType, false, TracerProviderConfig{
				CollectorURL:  tc.collectorUrl,
				ConsoleWriter: output,
			})
			assert.NoError(t, err)
			tracer := tp.Tracer(tc.name)

			_, span := tracer.Start(ctx, tc.name+"_normal")
			span.SetAttributes(attribute.String("testAttr", "testValue"))
			time.Sleep(time.Millisecond * 10)
			span.End()

			_, span = tracer.Start(ctx, tc.name+"_error")
			time.Sleep(time.Millisecond * 10)
			span.RecordErrorAndSetStatus(errors.New("some fake error"))
			span.End()

			// Shutdown the tracer to flush the output
			tp.Shutdown(ctx)

			tc.assertFunc(t, tracer, tp, output)
		})
	}

	t.Run("Errors on invalid provider", func(t *testing.T) {
		_, err := NewTracerProvider(TracerProviderType(10), false, TracerProviderConfig{})
		assert.Contains(t, err.Error(), "invalid TracerProviderType")
	})
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		description string
		pType       TracerProviderType
		config      TracerProviderConfig
		factory     func(TracerProviderType, bool, TracerProviderConfig) (TracerProvider, error)
		assertions  func(*testing.T, context.Context, Tracer, func(context.Context), error)
	}{
		{
			description: "failing factory returns error",
			factory: func(tpt TracerProviderType, b bool, tpc TracerProviderConfig) (TracerProvider, error) {
				return nil, fmt.Errorf("some error")
			},
			assertions: func(t *testing.T, ctx context.Context, tracer Tracer, f func(context.Context), err error) {
				assert.Nil(t, ctx)
				assert.Nil(t, tracer)
				assert.Nil(t, f)
				assert.Contains(t, err.Error(), "some error")
			},
		},
		{
			description: "noop tracer",
			pType:       NoopProviderType,
			config: TracerProviderConfig{
				TracerName:        "mockConfig",
				TracerService:     "mockService",
				TracerEnvironment: "mockEnv",
				TracerID:          1001,
				CollectorURL:      "https://mockUrl",
				ConsoleWriter:     os.Stdout,
			},
			factory: func(tpt TracerProviderType, b bool, tpc TracerProviderConfig) (TracerProvider, error) {
				return NewTracerProvider(tpt, b, tpc)
			},
			assertions: func(t *testing.T, ctx context.Context, tracer Tracer, f func(context.Context), err error) {
				var span Span
				ctx, span = tracer.Start(ctx, "some span")

				assert.Nil(t, err)
				assert.False(t, span.IsRecording())

				f(ctx)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctx, tracer, cleanup, err := create(
				tc.pType, tc.config, tc.factory,
			)
			tc.assertions(t, ctx, tracer, cleanup, err)
		})
	}
}
