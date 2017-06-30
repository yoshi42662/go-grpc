# go-grpc

Dockerized golang API application using grpc.

### Usage

- Get repository
```bash
$ git clone git@github.com:yoshi42662/go-grpc.git
```

- Build compose containers
```bash
$ docker-compose build
```

- Get into server container and run server.go
```bash
$ docker-compose run server bash
# -> Get into server container
(server) $ go run server/server.go
```

- Get into server container from different console and run client.go
```bash
# List all running containers
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
a69220a48eda        gogrpc_server       "bash"              20 minutes ago      Up 20 minutes       8080/tcp            gogrpc_server_run_1

# Get into gogrpc_server container with container ID.
$ docker exec -it a69220a48eda bash
# -> Get into the same server container
(server) $ go run client/client.go
```
