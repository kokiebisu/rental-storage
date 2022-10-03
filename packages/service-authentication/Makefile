.PHONY: prepare build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/sign_in sign_in.go generate_token.go &&
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/sign_up sign_up.go generate_token.go

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build