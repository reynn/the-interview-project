
build-client:
	go build -ldflags="-s -w" -trimpath -o client ./client-go/cmd/client/client.go

build-service:
	go build -ldflags="-s -w" -trimpath -o service ./service-go/cmd/service/service.go

build: build-client build-service

docker-client:
	docker build -t client ./client-go

docker-service:
	docker build -t service ./service-go

docker: docker-client docker-service

setup:
	./setup.sh

clean:
	./clean.sh
	rm -f client service

test-client:
	go test -v -race ./client-go/...

test-service:
	go test -v -race ./service-go/...

test: test-client test-service
