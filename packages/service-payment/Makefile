.PHONY: build
build:
	cd cmd/make_payment && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/make_payment && cd ../..
	cd cmd/create_payment_customer && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/create_payment_customer && cd ../..

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build

.PHONY: unit-test
unit-test:
	cd pkg/test && go test ./... && cd ../..
