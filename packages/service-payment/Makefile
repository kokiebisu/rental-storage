build:
	cd cmd/make_payment && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/make_payment *.go && cd ../.. &&
	cd cmd/create_payment_customer && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/create_payment_customer *.go && cd ../..

clean:
	rm -rf ./bin Gopkg.lock

prepare: clean build

.PHONY: test
test:
	cd cmd/create_payment_customer && go test && cd ../../
	cd cmd/make_payment && go test