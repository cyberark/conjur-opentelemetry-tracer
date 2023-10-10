# conjur-opentelemetry-tracer
A companion library for OpenTelemetry to provide tracing support for Conjur components
in a consistent fashion. Defines output providers for Jaeger, Console, and Noop.

## Certification level
TODO: Select the appropriate certification level section below, and remove all others.

{Community}
![](https://img.shields.io/badge/Certification%20Level-Community-28A745?link=https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md)

This repo is a **Community** level project. It's a community contributed project that **is not reviewed or supported
by CyberArk**. For more detailed information on our certification levels, see [our community guidelines](https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md#community).

{Trusted}
![](https://img.shields.io/badge/Certification%20Level-Trusted-007BFF?link=https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md)

This repo is a **Trusted** level project. It's been reviewed by CyberArk to verify that it will securely
work with Conjur Open Source as documented. For more detailed  information on our certification levels, see
[our community guidelines](https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md#community).

{Certified}
![](https://img.shields.io/badge/Certification%20Level-Certified-6C757D?link=https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md)

This repo is a **Certified** level project. It's been reviewed by CyberArk to verify that it will securely
work with CyberArk Conjur Enterprise as documented. In addition, CyberArk offers Enterprise-level support for these features. For
more detailed  information on our certification levels, see [our community guidelines](https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md#community).

## Requirements

TODO: Add any requirements that apply to your project here. Which Conjur Open Source / Enterprise versions is it
compatible with? Does it integrate with other tools / projects - and if so, what versions of those
does it require?

## Usage instructions

```go
// Get TracerProviderType from configuration
traceType, collectorUrl := trace.TypeFromEnv()

// Create a new Tracer
ctx, tracer, cleanup, _ := trace.Create(
    traceType,
    TracerProviderConfig{
        TracerName:        "my-tracer",
        TracerService:     "tracer-service",
        TracerEnvironment: "development",
        TracerID:          1,
        CollectorURL:      collectorUrl,
        ConsoleWriter:     os.Stdout,
    },
)
defer cleanup(ctx)

// Setup a Span
ctx, span := tracer.Start(ctx, "My process")
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

Copyright (c) 2022 CyberArk Software Ltd. All rights reserved.

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
