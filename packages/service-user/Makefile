NAMESPACE = @service-user

.PHONY: set-profile
set-profile:
	@echo "Setting AWS Profile..."
	@export AWS_PROFILE=rental-storage

.PHONY: build
build:
	@echo "[ $(NAMESPACE) ] Building..."
	@cd cmd/controller/create-user && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-user && cd ../..
	@cd cmd/controller/find-user-by-email && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-user-by-email && cd ../..
	@cd cmd/controller/find-user-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-user-by-id && cd ../..
	@cd cmd/controller/remove-user-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/remove-user-by-id && cd ../..

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
	@echo "[ $(NAMESPACE) ] Running Unit Tests..."
	@cd test/unit && export ENV=test && go test ./... && cd ../..

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running Integration Tests..."
	@cd test/integration && export ENV=test && go test ./... && cd ../..

.PHONY: lint
lint:
	@echo "[ $(NAMESPACE) ] Linting..."
	@golangci-lint run ./... --timeout=5m

.PHONY: mock-generate
mock-generate:
	@echo "[ $(NAMESPACE) ] Generating Mocks..."
	@mockery --all --case snake


.PHONY: deps-start
deps-start:
	@echo "[$(NAMESPACE) ] Running Dependencies..."
	@docker-compose up -d --build

.PHONY: deps-stop
deps-stop:
	@echo "[$(NAMESPACE) ] Stopping Dependencies..."
	@docker-compose down --volumes --remove-orphans
