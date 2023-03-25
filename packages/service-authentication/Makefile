NAMESPACE = @service-authentication

.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage..."
	@export AWS_PROFILE=rental-storage
	
.PHONY: build
build:
	@echo "[ $(NAMESPACE) ] Building..."
	@cd cmd/controller/signin && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/signin && cd ../..
	@cd cmd/controller/signup && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/signup && cd ../..
	@cd cmd/controller/verify && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/verify && cd ../..

.PHONY: clean
clean:
	@echo "[ $(NAMESPACE) ] Cleaning..."
	@rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: 
	@echo "[ $(NAMESPACE) ] Preparing..."
	@clean build

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running integration tests..."
	@export ENV=test && cd test/integration && go test ./... -v && cd ../..

.PHONY: test-unit
test-unit:
	@echo "[ $(NAMESPACE) ] Running unit tests..."
	@export ENV=test && cd test/unit && go test ./... -v && cd ../..

.PHONY: lint
lint:
	@echo "[ $(NAMESPACE) ] Linting..."
	@golangci-lint run ./... --timeout=5m

.PHONY: mock-generate
mock-generate:
	@echo "[ $(NAMESPACE) ] Generating mocks..."
	@mockery --all --case snake

.PHONY: deps-start
deps-start:
	@echo "[ $(NAMESPACE) ] Running dependencies..."
	@docker-compose up -d --build

.PHONY: deps-end
deps-stop:
	@echo "[ $(NAMESPACE) ] Stopping dependencies..."
	@docker-compose down --volumes
