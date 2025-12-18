# conjur-opentelemetry-tracer
A companion library for OpenTelemetry to provide tracing support for Conjur components
in a consistent fashion. Defines output providers for Jaeger, Console, and Noop.

## Certification level

![](https://img.shields.io/badge/Certification%20Level-Community-28A745?link=https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md)

This repo is a **Community** level project. It's a community contributed project that **is not reviewed or supported
by CyberArk**. For more detailed information on our certification levels, see [our community guidelines](https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md#community).

## Usage instructions

```go

tp, _ := trace.NewTracerProvider(trace.ConsoleProviderType, "", os.Stdout, true)
defer tp.Shutdown()

tracer := tp.Tracer("my-service")

ctx, span := tracer.Start(context.Background(), "My process")
defer span.End()

// Do some task

if err != nil {
   span.RecordErrorAndSetStatus(err)
}

```

## Contributing

We welcome contributions of all kinds to this repository. For instructions on how to get started and descriptions
of our development workflows, please see our [contributing guide](CONTRIBUTING.md).

## License

Copyright (c) 2025 CyberArk Software Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

For the full license text see [`LICENSE`](LICENSE).
