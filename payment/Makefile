.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/payment.proto

.PHONY: build
build:
	@echo Building for Linux...
    GOOS=linux GOARCH=amd64 go build -o order-service *.go
.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t order-service:latest