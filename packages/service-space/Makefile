NAMESPACE = @service-space

.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage..."
	@export AWS_PROFILE=rental-storage

.PHONY: clean
clean:
	@echo "[ $(NAMESPACE) ] Cleaning up..."
	@rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo

.PHONY: build
build:
	@echo "[ $(NAMESPACE) ] Building..."
	@cd cmd/controller/find-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-space && cd ../..
	@cd cmd/controller/find-spaces && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-spaces && cd ../..
	@cd cmd/controller/create-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-space && cd ../..
	@cd cmd/controller/delete-space && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/delete-space && cd ../..

.PHONY: test-unit
test-unit:
	@echo "[ $(NAMESPACE) ] Running unit tests..."
	@cd test/unit && export ENV=test && go test ./... && cd ../..

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running integration tests..."
	@cd test/integration && export ENV=test && go test ./... && cd ../..

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
	@docker-compose down --volumes --remove-orphans
