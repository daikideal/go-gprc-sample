# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine

WORKDIR /go-grpc-sample

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go/ .

RUN go build -o ./bin/go-grpc-sample
