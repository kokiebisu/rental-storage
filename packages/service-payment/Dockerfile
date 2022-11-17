FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN cd cmd/controller/make-payment && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/make-payment
RUN cd cmd/controller/create-payment-customer && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-payment-customer

FROM amazon/aws-lambda-go:latest

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]