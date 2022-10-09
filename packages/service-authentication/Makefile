.PHONY: prepare build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/sign_in sign_in.go generate_jwt_token.go &&
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/sign_up sign_up.go generate_jwt_token.go &&
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/verify verify.go verify_token.go generate_jwt_token.go

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build