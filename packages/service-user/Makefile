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

.PHONY: test-unit
test-unit:
	cd test/unit && export ENV=test && go test ./... && cd ../..

.PHONY: test-integration
test-integration:
	cd test/integration && export ENV=test && go test ./... && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m

.PHONY: mock-generate
mock-generate:
	mockery --all --case snake


.PHONY: deps-start
deps-start:
	docker-compose up -d --build

.PHONY: deps-end
deps-stop:
	docker-compose down --volumes

.PHONY: remove
remove:
	npx -y serverless@3.28.1 remove --config serverless.yml --stage local

.PHONY: deploy
deploy:
	npx -y serverless@3.28.1 deploy --config serverless.yml --stage local
