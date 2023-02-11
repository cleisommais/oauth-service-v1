FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# Install golangci
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0

# Installing buffalo binary
COPY buffalo /go/bin/buffalo

WORKDIR /
RUN go install github.com/gobuffalo/buffalo-pop/v3@latest

RUN mkdir /src
WORKDIR /src
