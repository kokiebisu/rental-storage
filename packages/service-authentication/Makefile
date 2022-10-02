.PHONY: prepare build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/signin signin.go && cd ..

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build