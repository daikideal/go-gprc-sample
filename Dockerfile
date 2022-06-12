# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine AS build

WORKDIR /go-grpc-sample

RUN apk update && apk add unzip curl

# protoc installation
ENV PROTOC_VERSION=21.1
ENV PROTOC_ZIP_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN mkdir -p /tmp && \
    curl -L -o /tmp/protoc.zip ${PROTOC_ZIP_URL} && \
    unzip /tmp/protoc.zip -d /tmp/protoc && \
    cp /tmp/protoc/bin/protoc /usr/local/bin/protoc && \
    rm -rf /tmp/protoc*

# go plugins for protoc installation (`go get` is deprecated since go 1.17)
ENV PROTOC_GEN_VERSION=1.28
ENV PROTOC_GEN_GO_GRPC_VERSION=1.2

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_VERSION} && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go/ .

RUN go build -o ./bin/go-grpc-sample

FROM ubuntu:22.10 AS run

WORKDIR /root

COPY --from=build /go-grpc-sample/bin/go-grpc-sample ./bin/go-grpc-sample

CMD [ "./bin/go-grpc-sample" ]
