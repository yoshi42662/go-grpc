# go-grpc

Dockerized golang API application using gRPC.

### Quick Start

- Get repository
```bash
$ git clone git@github.com:yoshi42662/go-grpc.git
```

- Build compose containers
```bash
$ docker-compose build
```

- Run
```bash
$ docker-compose up
```

### Generate .go source from your own .proto file.

- Get into server container
```bash
$ docker-compose run server bash
```

- Run `protoc` command
```bash
root@container_id:/go/src/server# protoc --go_out=plugins=grpc:. pb/*.proto
```
