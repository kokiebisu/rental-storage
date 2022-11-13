FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN cd cmd/make_payment && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/make_payment
RUN cd cmd/create_payment_customer && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../bin/create_payment_customer

FROM amazon/aws-lambda-go:latest

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]