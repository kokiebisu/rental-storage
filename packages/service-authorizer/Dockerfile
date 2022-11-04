FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN cd cmd/authorize && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../../bin/authorize *.go

FROM amazon/aws-lambda-go:latest

WORKDIR ${LAMBDA_TASK_ROOT}

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]
