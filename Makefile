PROJECT_NAME ?= go_app
CLI ?= @docker exec -it ${PROJECT_NAME}_cli
VOLUME ?= $(shell pwd)

docker.start.all:
	docker-compose --project-name $(PROJECT_NAME) -f ./deployments/docker-compose/docker-compose.yml up -d

docker.stop.all:
	docker-compose --project-name $(PROJECT_NAME) -f ./deployments/docker-compose/docker-compose.yml stop

docker.restart.all: docker.stop.all docker.start.all

cli:
	${CLI} /bin/sh

proto.generate:
	${CLI} protoc --go_out=plugins=grpc:. api/proto/*.proto