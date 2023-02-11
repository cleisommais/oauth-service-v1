FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# Installing buffalo binary
COPY buffalo /go/bin/buffalo

WORKDIR /
RUN go install github.com/gobuffalo/buffalo-pop/v3@latest

RUN mkdir /src
WORKDIR /src
