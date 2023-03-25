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

.PHONY: test-unit
test-unit:
	@echo "[ $(NAMESPACE) ] Running unit tests"
	@cd test && go test ./... && cd ../..

.PHONY: lint
lint:
	@echo "[ $(NAMESPACE) ] Linting..."
	@golangci-lint run ./... --timeout=5m
