.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage"
	@export AWS_PROFILE=rental-storage
	
.PHONY: build
build:
	@echo "Building..."
	@cd cmd/controller/signin && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/signin && cd ../..
	@cd cmd/controller/signup && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/signup && cd ../..
	@cd cmd/controller/verify && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/verify && cd ../..

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: 
	@echo "Preparing..."
	@clean build

.PHONY: test-integration
test-integration:
	@echo "Running integration tests"
	@export ENV=test && cd test/integration && go test ./... -v && cd ../..

.PHONY: test-unit
test-unit:
	@echo "Running unit tests"
	@export ENV=test && cd test/unit && go test ./... -v && cd ../..

.PHONY: lint
lint:
	@echo "Linting"
	@golangci-lint run ./... --timeout=5m

.PHONY: mock-generate
mock-generate:
	@echo "Generating mocks"
	@mockery --all --case snake

.PHONY: deps-start
deps-start:
	@echo "Running Mock Server..."
	@docker-compose up -d --build

.PHONY: deps-end
deps-stop:
	@echo "Stopping Mock Server..."
	@docker-compose down --volumes
