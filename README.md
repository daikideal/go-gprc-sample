# go-gprc-sample

## build

Clone repository.

```bash
$ git clone git@github.com:daikideal/go-grpc-sample
```

Build docker image.

```bash
$ docker build -t go-grpc-sample .
```

## run

Run docker container and enter the container.

```bash
$ docker run -it \
  -v $PWD:/go/src/github.com/daikideal/go-grpc-sample \
  go-grpc-sample \
  bash
```

Start go server (background).

```bash
# go run cmd/server/server.go &
```

Send request with go cli.

```bash
# go run cmd/client/client.go
```
