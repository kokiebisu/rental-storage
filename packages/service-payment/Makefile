.PHONY: build
build:
	cd cmd/controller/make-payment && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/make-payment && cd ../..
	cd cmd/controller/create-payment-customer && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-payment-customer && cd ../..

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build

.PHONY: test-unit
test-unit:
	cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m
