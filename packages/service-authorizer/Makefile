.PHONY: prepare build clean

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/handler authorizer.go

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build
