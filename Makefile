IMAGE = eth-caching-proxy

.PHONY: all
all: build run

.PHONY: build
build:
	go build -o bin/main main.go

.PHONY: run
run:
	go run main.go

.PHONY: docker_build
docker_build:
	docker build -t ${IMAGE} .

.PHONY: docker_run
docker_run:
	docker run -p 9000:9000 ${IMAGE}