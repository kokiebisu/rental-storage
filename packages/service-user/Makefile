.PHONY: build
build:
	cd cmd/controller/create_user && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create_user && cd ../..
	cd cmd/controller/find_user_by_email && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find_user_by_email && cd ../..
	cd cmd/controller/find_user_by_id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find_user_by_id && cd ../..
	cd cmd/controller/remove_user_by_id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/remove_user_by_id && cd ../..

.PHONY: clean
clean:
	rm -rf ./bin Gopkg.lock

.PHONY: prepare
prepare: clean build

.PHONY: unit-test
unit-test:
	cd test && go test ./... && cd ../..
