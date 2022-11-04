.PHONY: build
build:
	cd cmd/authorize && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/authorize authorizer.go

.PHONY: clean
clean:
	rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: 
	clean build

.PHONY: unit-test
unit-test: 
	cd cmd/authorize && go test ./.../unit_test

.PHONY: integration-test
integration-test: 
	go test ./.../integration_test