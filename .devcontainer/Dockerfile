FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# Install golangci
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0

RUN wget https://github.com/gobuffalo/cli/releases/download/v0.18.13/buffalo_0.18.13_Linux_x86_64.tar.gz \
  && tar -xvzf buffalo_0.18.13_Linux_x86_64.tar.gz \
  && sudo mv buffalo /usr/local/bin/buffalo
