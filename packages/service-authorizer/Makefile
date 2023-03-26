NAMESPACE = @service-authorizer

.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage"
	@export AWS_PROFILE=rental-storage
	
.PHONY: build
build:
	@echo "[ $(NAMESPACE) ] Building..."
	@cd cmd/controller/authorize && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../../bin/authorize *.go

.PHONY: clean
clean:
	@echo "[ $(NAMESPACE) ] Cleaning..."
	@rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: 
	@echo "[ $(NAMESPACE) ] Preparing..."
	@clean build

.PHONY: mock-generate
mock-generate:
	@echo "[ $(NAMESPACE) ] Generating mocks..."
	@mockery --all --case snake

.PHONY: test-unit
test-unit:
	@echo "[ $(NAMESPACE) ] Running unit tests..."
	@cd test/unit && go test ./... && cd ../..

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running integration tests..."
	@cd test/integration && go test ./... && cd ../..

.PHONY: lint
lint:
	@echo "[ $(NAMESPACE) ] Linting..."
	@golangci-lint run ./... --timeout=5m
