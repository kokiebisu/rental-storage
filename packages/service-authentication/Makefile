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
	cd pkg/test && go test ./... && cd ../..
