arg = $(filter-out $@,$(MAKECMDGOALS))
this_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

proto:
	@protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. \app/grpcHandler/fiber.proto

proto_s:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/grpcHandler/fiber.proto

run:
	@go run main.go $(call arg)

build:
	@go build

run-nohup: build
	@nohup ./go-fiber-v2 $(call arg) &

docker:
	@docker build . -t go-fiber-v2
	@docker image prune --filter label=stage-go-fiber-v2=builder -f

docker-run:
	docker run -v $(this_dir)logs:/app/logs -e "environment=$(arg)" -d --restart always --hostname go-fiber-v2 --name go-fiber-v2 -p 9098:8888 -e TZ=Asia/Jakarta --link go-otto-users:go-otto-users --link redis_fiber:redis go-fiber-v2

docker-stop:
	@docker stop go-fiber-v2-v2

clear-container:
	@docker rm -f go-fiber-v2

docker-stop-rm: docker-stop clear-container

clear-image:
	@docker rmi -f go-fiber-v2

clear-docker: clear-container clear-image

docker-exec:
	@docker exec -it go-fiber-v2 sh

volume:
	@docker volume create $(call arg)

run-redis:
	@docker run --detach --restart always --name redis_fiber --hostname redis.fiber redis redis-server

run-redis-volume:
	@docker run --detach --restart always -v $(call arg):/data --name redis_fiber --hostname redis.fiber redis redis-server

redis-cli:
	@docker run -it  --link redis_fiber:redis --rm redis redis-cli -h redis -p 6379

stop-redis:
	@docker stop redis_fiber
	@docker rm redis_fiber