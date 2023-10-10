package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeFromAnnotationsOrEnv(t *testing.T) {
	testCases := []struct {
		description string
		env         map[string]string
		annots      map[string]string
		assertions  func(*testing.T, TracerProviderType, string)
	}{
		{
			description: "annotations override envvars",
			env: map[string]string{
				"JAEGER_COLLECTOR_URL": "http://collector",
			},
			annots: map[string]string{
				"conjur.org/log-traces": "true",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Empty(t, s)
				assert.Equal(t, ConsoleProviderType, int(tpt))
			},
		},
		{
			description: "envvars used if annotations absent",
			env: map[string]string{
				"JAEGER_COLLECTOR_URL": "http://collector",
			},
			annots: map[string]string{},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Equal(t, "http://collector", s)
				assert.Equal(t, JaegerProviderType, int(tpt))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pType, collectorUrl := typeFromAnnotationsOrCustomEnv(
				tc.annots,
				func(key string) string {
					return tc.env[key]
				},
			)
			tc.assertions(t, pType, collectorUrl)
		})
	}
}

func TestTypeFromEnv(t *testing.T) {
	testCases := []struct {
		description string
		env         map[string]string
		assertions  func(*testing.T, TracerProviderType, string)
	}{
		{
			description: "noop type",
			env:         map[string]string{},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Empty(t, s)
				assert.Equal(t, NoopProviderType, int(tpt))
			},
		},
		{
			description: "console type",
			env: map[string]string{
				"LOG_TRACES": "true",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Empty(t, s)
				assert.Equal(t, ConsoleProviderType, int(tpt))
			},
		},
		{
			description: "jaeger type",
			env: map[string]string{
				"JAEGER_COLLECTOR_URL": "http://collector",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Equal(t, "http://collector", s)
				assert.Equal(t, JaegerProviderType, int(tpt))
			},
		},
		{
			description: "jaeger type overrides console type",
			env: map[string]string{
				"LOG_TRACES":           "true",
				"JAEGER_COLLECTOR_URL": "http://collector",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Equal(t, "http://collector", s)
				assert.Equal(t, JaegerProviderType, int(tpt))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pType, collectorUrl := typeFromCustomEnv(
				func(key string) string {
					return tc.env[key]
				},
			)
			tc.assertions(t, pType, collectorUrl)
		})
	}
}

func TestTypeFromAnnotations(t *testing.T) {
	testCases := []struct {
		description string
		annots      map[string]string
		assertions  func(*testing.T, TracerProviderType, string)
	}{
		{
			description: "noop type",
			annots:      map[string]string{},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Empty(t, s)
				assert.Equal(t, NoopProviderType, int(tpt))
			},
		},
		{
			description: "console type",
			annots: map[string]string{
				"conjur.org/log-traces": "true",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Empty(t, s)
				assert.Equal(t, ConsoleProviderType, int(tpt))
			},
		},
		{
			description: "jaeger type",
			annots: map[string]string{
				"conjur.org/jaeger-collector-url": "http://collector",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Equal(t, "http://collector", s)
				assert.Equal(t, JaegerProviderType, int(tpt))
			},
		},
		{
			description: "jaeger type overrides console type",
			annots: map[string]string{
				"conjur.org/log-traces":           "true",
				"conjur.org/jaeger-collector-url": "http://collector",
			},
			assertions: func(t *testing.T, tpt TracerProviderType, s string) {
				assert.Equal(t, "http://collector", s)
				assert.Equal(t, JaegerProviderType, int(tpt))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pType, collectorUrl := TypeFromAnnotations(tc.annots)
			tc.assertions(t, pType, collectorUrl)
		})
	}
}
