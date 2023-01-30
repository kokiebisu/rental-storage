.PHONY: build
build:
	cd cmd/controller/create-user && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-user && cd ../..
	cd cmd/controller/find-user-by-email && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-user-by-email && cd ../..
	cd cmd/controller/find-user-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-user-by-id && cd ../..
	cd cmd/controller/remove-user-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/remove-user-by-id && cd ../..

.PHONY: clean
clean:
	rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: clean build

.PHONY: test
test:
	cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: generate-mock
generate-mock:
	mockery --all