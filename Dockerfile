# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine AS build

WORKDIR /go-grpc-sample

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go/ .

RUN go build -o ./bin/go-grpc-sample

FROM ubuntu:22.10 AS run

WORKDIR /root

COPY --from=build /go-grpc-sample/bin/go-grpc-sample ./bin/go-grpc-sample

CMD [ "./bin/go-grpc-sample" ]
