# syntax=docker/dockerfile:1

FROM golang:1.18.3

WORKDIR /go/src/github.com/daikideal/go-grpc-sample

RUN apt update && apt install unzip

# protoc installation
ARG PROTOC_VERSION=21.1
ARG PROTOC_ZIP_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN mkdir -p /tmp && \
    cd /tmp && \
    curl -L -o ./protoc.zip ${PROTOC_ZIP_URL} && \
    unzip ./protoc.zip -d ./protoc && \
    cp ./protoc/bin/protoc /usr/local/bin/protoc && \
    rm -rf ./protoc*

# go plugins for protoc installation (`go get` is deprecated since go 1.17)
ARG PROTOC_GEN_VERSION=1.28
ARG PROTOC_GEN_GO_GRPC_VERSION=1.2

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_VERSION} && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
