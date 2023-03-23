.PHONY: profile-set
profile-set:
	export AWS_PROFILE=rental-storage

.PHONY: test-env-set
test-env-set:
	export ENV=test

.PHONY: build
build:
	make profile-set
	cd cmd/controller/accept-booking && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/accept-booking && cd ../..
	cd cmd/controller/create-booking && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-booking && cd ../..
	cd cmd/controller/find-booking-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-booking-by-id && cd ../..
	cd cmd/controller/find-bookings && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-bookings && cd ../..

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
	make profile-set
	make test-env-start
	cd test/integration && go test -count=1 ./... && cd ../..

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

.PHONY: aws-docker-check-items
aws-docker-check-items: 
	make profile-set
	aws dynamodb scan --table-name booking --endpoint-url http://localhost:8000
