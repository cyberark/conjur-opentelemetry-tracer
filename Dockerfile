#=============== Build Container ===================
FROM golang:1.23

ARG GIT_COMMIT_SHORT="dev"
ARG KUBECTL_VERSION=1.31.2

# On CyberArk dev laptops, golang module dependencies are downloaded with a
# corporate proxy in the middle. For these connections to succeed we need to
# configure the proxy CA certificate in build containers.
#
# To allow this script to also work on non-CyberArk laptops where the CA
# certificate is not available, we copy the (potentially empty) directory
# and update container certificates based on that, rather than rely on the
# CA file itself.
COPY build_ca_certificate /usr/local/share/ca-certificates/
RUN update-ca-certificates

RUN mkdir -p /work
WORKDIR /work

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

# source files
COPY pkg ./pkg

# The `gitCommitShort` override is there to provide the git commit information in the final
# binary.
RUN go build \
    -ldflags="-X github.com/cyberark/conjur-opentelemetry-tracer/pkg/version.gitCommitShort=$GIT_COMMIT_SHORT" \
    -o cyberark-conjur-opentelemetry-tracer \
    ./pkg/trace
