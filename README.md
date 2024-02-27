# go-fiber-v2

Skeleton service using golang and fiber framework. Mock example available.

## Prepare service

```
$ go mod tidy
```

## GRPC Doc

Generate: https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file

buf gen: https://github.com/bufbuild/buf

Other https://protobuf.dev/reference/go/go-generated/#package 


## How To Run

env mode available: local/dev/prod

#### run

```sh
$ make run
```

#### Nohup

```sh
$ make run-nohup {env}
```

#### Docker (create redis first if needed, instructions at the bottom)

- Build Docker Image and Delete Stage Builder

```sh
$ make docker
```

- Add and Start Container

```sh
$ make docker-run {env}
```

- Stop Container

```sh
$ make docker-stop
```

- Stop and remove Container

```sh
$ make docker-stop-rm
```

- Delete Container

```sh
$ make clear-container
```

- Clear Docker

```sh
$ make clear-docker
```

- Docker Exec

```sh
$ make docker-exec
```

## Create Redis in Docker if needed

#### 1. Create Volume (optional)

```sh
$ make volume {volume_name_without_space}
```

#### 2. Add Container and Start Redis

- with volume

```sh
$ make run-redis-volume {volume_name}
```

- without volume

```sh
$ make run-redis
```

#### 3. Redis CLI

```sh
$ make redis-cli
```

#### 4. Stop Redis

```sh
$ make stop-redis
```

## Go Unit Test

```sh
$ go test ./... -v
```

## Go Unit Test Coverage

```sh
$ go test -v -coverpkg=./... -coverprofile=coverage.out ./...
$ go tool cover -func=coverage.out
```

## Go Security

install gosec: brew install gosec

```sh
$ gosec ./...
```