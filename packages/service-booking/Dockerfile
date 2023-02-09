FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN cd cmd/controller/create-booking && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/create-booking
RUN cd cmd/controller/find-booking-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-booking-by-id

FROM amazon/aws-lambda-go:latest

WORKDIR ${LAMBDA_TASK_ROOT}

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]