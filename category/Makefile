.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/category.proto

.PHONY: build
build:
	go build -o order-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t order-service:latest