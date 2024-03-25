module github.com/cyberark/conjur-opentelemetry-tracer

go 1.22

require (
	github.com/stretchr/testify v1.7.2
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger v1.7.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c => golang.org/x/sys v0.8.0

replace golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 => golang.org/x/sys v0.8.0

replace golang.org/x/sys v0.1.0 => golang.org/x/sys v0.8.0
