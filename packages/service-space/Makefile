.PHONY: clean
clean:
	rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo

.PHONY: build
build:
	cd cmd/controller/find-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-space && cd ../..
	cd cmd/controller/find-spaces && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-spaces && cd ../..
	cd cmd/controller/create-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-space && cd ../..
	cd cmd/controller/delete-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/delete-space && cd ../..

.PHONY: test
test:
	cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m

.PHONY: generate-mock
generate-mock:
	mockery --all --case snake
