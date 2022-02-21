PROJECT_NAME=go_app

docker.start.all:
	docker-compose --project-name $(PROJECT_NAME) -f ./deployments/docker-compose/docker-compose.yml up -d

docker.stop.all:
	docker-compose --project-name $(PROJECT_NAME) -f ./deployments/docker-compose/docker-compose.yml stop

docker.restart.all: docker.stop.all docker.start.all