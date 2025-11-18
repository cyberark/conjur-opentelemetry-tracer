# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.4] - 2025-10-30

### Security
- Update Go to 1.25 (CONJSE-2067)

### Added
- Added `close-stale.yml` GitHub workflow

## [0.0.3] - 2025-10-16

### Changed
- Updated documentation to align with Conjur Enterprise name change to Secrets Manager. (CNJR-10997)


## [0.0.2] - 2024-11-04

### Security
- Update Go to 1.23 (CONJSE-1842, CNJR-6449)
- Update Go to 1.20 and golang.org/x/sys to v0.8.0
  [cyberark/conjur-opentelemetry-tracer#9](https://github.com/cyberark/conjur-opentelemetry-tracer/pull/9)

## [0.0.1] - 2022-01-12

### Added
- Basic functionality [cyberark/conjur-opentelemetry-tracer#1](https://github.com/cyberark/conjur-opentelemetry-tracer/pull/1)

### Changed
- Added replace for gopkg.in/yaml.v3 to ensure we use latest version in dep tree
  [cyberark/conjur-opentelemetry-tracer#6](https://github.com/cyberark/conjur-opentelemetry-tracer/pull/6)
- Updated go dependencies to latest versions (github.com/stretchr/testify -> 1.7.2, 
  go.opentelemetry.io/otel/* -> 1.7.0)
  [cyberark/conjur-opentelemetry-tracer#5](https://github.com/cyberark/conjur-opentelemetry-tracer/pull/5)

### Security
- Update golang.org/x/sys to 0.1.0 for CVE-2022-29526 (not vulnerable)
  [cyberark/conjur-opentelemetry-tracer#8](https://github.com/cyberark/conjur-opentelemetry-tracer/pull/8)

[Unreleased]: https://github.com/cyberark/secrets-provider-for-k8s/compare/v0.0.4...HEAD
[0.0.4]: https://github.com/cyberark/secrets-provider-for-k8s/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/cyberark/secrets-provider-for-k8s/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/cyberark/secrets-provider-for-k8s/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/cyberark/secrets-provider-for-k8s/releases/tag/v0.0.1
