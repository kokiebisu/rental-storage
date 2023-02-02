.PHONY: clean
clean:
	rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo

.PHONY: build
build:
	cd cmd/controller/find-listing-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-listing-by-id && cd ../..
	cd cmd/controller/find-listings-within-lat-lng && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-listings-within-lat-lng && cd ../..
	cd cmd/controller/add-listing && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/add-listing && cd ../..

.PHONY: test
test:
	cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: generate-mock
generate-mock:
	mockery --all --case snake
