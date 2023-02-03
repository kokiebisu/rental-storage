.PHONY: build
build:
	cd cmd/authorize && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/authorize *.go

.PHONY: clean
clean:
	rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: 
	clean build

.PHONY: test
test:
	cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m
