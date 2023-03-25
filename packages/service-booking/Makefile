NAMESPACE = @service-booking

.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage..."
	@export AWS_PROFILE=rental-storage

.PHONY: build
build:
	@echo "[ $(NAMESPACE) ] Building..."
	@cd cmd/controller/accept-booking && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/accept-booking && cd ../..
	@cd cmd/controller/create-booking && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-booking && cd ../..
	@cd cmd/controller/find-booking-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-booking-by-id && cd ../..
	@cd cmd/controller/find-bookings && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-bookings && cd ../..

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
	@echo "[ $(NAMESPACE) ] Running unit tests..."
	@cd test/unit && export ENV=test && go test ./... && cd ../..

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running integration tests..."
	@export AWS_PROFILE=rental-storage && export ENV=test && cd test/integration && go test -count=1 ./... && cd ../..

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

.PHONY: aws-docker-check-items
aws-docker-check-items:
	@echo "[ $(NAMESPACE) ] Checking items..."
	@make set-profile
	@aws dynamodb scan --table-name booking --endpoint-url http://localhost:8000
