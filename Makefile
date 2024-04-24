CLIENT_CONTAINER_NAME := client
SERVICE_CONTAINER_NAME := service

build-client:
	go build -ldflags="-s -w" -trimpath -o client ./client-go/cmd/client/client.go

build-service:
	go build -ldflags="-s -w" -trimpath -o service ./service-go/cmd/service/service.go

build: build-client build-service

docker-client:
	docker build -t $(CLIENT_CONTAINER_NAME) ./client-go

docker-service:
	docker build -t $(SERVICE_CONTAINER_NAME) ./service-go

docker: docker-client docker-service

setup:
	./setup.sh

clean:
	./clean.sh
	rm -f client service

run-service: docker-service
	docker run $(DOCKER_ARGS) --rm --env-file ./service-go/env/local.env $(SERVICE_CONTAINER_NAME)

test-client:
	go test -v -race ./client-go/...

test-service:
	go test -v -race ./service-go/...

test: test-client test-service
