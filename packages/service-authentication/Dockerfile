FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN cd cmd/signin && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/signin *.go
RUN cd cmd/signup && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/signup *.go
RUN cd cmd/verify && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/verify *.go

FROM amazon/aws-lambda-go:latest

WORKDIR ${LAMBDA_TASK_ROOT}

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]