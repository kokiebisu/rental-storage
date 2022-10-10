.PHONY: prepare build clean deploy

build:
	cd cmd/signin && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/signin *.go && cd ../.. &&
	cd cmd/signup && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/signup *.go && cd ../.. &&
	cd cmd/verify && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/verify *.go && cd ../..

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build