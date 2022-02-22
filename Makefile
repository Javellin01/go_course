PROJECT_NAME ?= go_app
CLI_CONTAINER ?= ${PROJECT_NAME}_cli
CLI ?= @docker exec -it ${CLI_CONTAINER}
VOLUME ?= $(shell pwd)

docker.start.all:
	docker-compose --project-name $(PROJECT_NAME) -f ./deployments/docker-compose/docker-compose.yml up -d

docker.stop.all:
	docker-compose --project-name $(PROJECT_NAME) -f ./deployments/docker-compose/docker-compose.yml stop

docker.restart.all: docker.stop.all docker.start.all

cli.run:
	@docker rm -f ${CLI_CONTAINER}
	@docker build -t ${PROJECT_NAME}:cli -f ./deployments/docker/cli/Dockerfile .
	@docker run -it -d --name=${CLI_CONTAINER} -v ${VOLUME}:/go/src/service ${PROJECT_NAME}:cli

proto.generate:
	${CLI} protoc --go_out=plugins=grpc:. api/proto/*.proto