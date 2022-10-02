.PHONY: prepare build clean deploy

build:
	cd handler && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../bin/handler authorizer.go && cd ..

clean:
	rm -rf ./bin ./vendor Gopkg.lock

prepare: clean build