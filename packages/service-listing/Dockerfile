FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN cd cmd/controller/find-listing-by-id && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-listing-by-id
RUN cd cmd/controller/find-listings-within-lat-lng && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/find-listings-within-lat-lng
RUN cd cmd/controller/add-listing && env GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags="-s -w" -o ../../../bin/add-listing

FROM amazon/aws-lambda-go:latest

WORKDIR ${LAMBDA_TASK_ROOT}

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]