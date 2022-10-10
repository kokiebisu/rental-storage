.PHONY: prepare build clean

build:
	cd cmd/make_payment && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/make_payment *.go && cd ../.. &&
	cd cmd/generate_customer_id && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/generate_customer_id *.go && cd ../..

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build
