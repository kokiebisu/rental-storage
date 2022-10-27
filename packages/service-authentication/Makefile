.PHONY: prepare build clean deploy

build:
	cd cmd/signin && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/signin *.go && cd ../..
	cd cmd/signup && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/signup *.go && cd ../..
	cd cmd/verify && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/verify *.go && cd ../..

clean:
	rm -rf ./bin Gopkg.lock

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