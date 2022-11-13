.PHONY: build
build:
	cd cmd/signin && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/signin && cd ../..
	cd cmd/signup && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/signup && cd ../..
	cd cmd/verify && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/verify && cd ../..

.PHONY: clean
clean:
	rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: clean build

.PHONY: unit-test
unit-test:
	cd cmd/signin && go test ./.../unit_test && cd ../..
	cd cmd/signup && go test ./.../unit_test && cd ../..
	cd cmd/verify && go test ./.../unit_test && cd ../..

.PHONY: integration-test
integration-test: 
	cd cmd/signin && go test ./.../integration_test && cd ../..
	cd cmd/signup && go test ./.../integration_test && cd ../..
	cd cmd/verify && go test ./.../integration_test && cd ../..