.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage"
	@export AWS_PROFILE=rental-storage
	
.PHONY: build
build:
	@echo "Building..."
	@cd cmd/controller/authorize && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../../bin/authorize *.go

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: 
	@echo "Preparing..."
	@clean build

.PHONY: test-unit
test-unit:
	@echo "Running unit tests"
	@cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	@echo "Linting"
	@golangci-lint run ./... --timeout=5m
