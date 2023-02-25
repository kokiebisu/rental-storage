.PHONY: clean
clean:
	rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo

.PHONY: build
build:
	cd cmd/controller/find-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-space && cd ../..
	cd cmd/controller/find-spaces && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-spaces && cd ../..
	cd cmd/controller/create-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-space && cd ../..
	cd cmd/controller/delete-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/delete-space && cd ../..

.PHONY: test-unit
test-unit:
	cd test/unit && go test ./... && cd ../..

.PHONY: test-integration
test-integration:
	cd test/integration && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m

.PHONY: generate-mock
generate-mock:
	mockery --all --case snake

.PHONY: deps-start
deps-start:
	docker-compose up -d --build

.PHONY: deps-end
deps-end:
	docker-compose down --volumes
