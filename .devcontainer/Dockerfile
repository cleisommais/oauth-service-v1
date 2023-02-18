FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# Install golangci
# https://golangci-lint.run/usage/install/#local-installation
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.1

RUN golangci-lint --version
