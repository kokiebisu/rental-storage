.PHONY: build
build:
	cd cmd/controller/signin && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/signin && cd ../..
	cd cmd/controller/signup && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/signup && cd ../..
	cd cmd/controller/verify && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/verify && cd ../..

.PHONY: clean
clean:
	rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: clean build

.PHONY: test-unit
test-unit:
	cd test && go test ./... -v && cd ../..

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m

.PHONY: generate-mock
generate-mock:
	mockery --all --case snake

.PHONY: remove
remove:
	npx -y serverless@3.28.1 remove --config serverless.yml --stage local

.PHONY: deploy
deploy:
	npx -y serverless@3.28.1 deploy --config serverless.yml --stage local
