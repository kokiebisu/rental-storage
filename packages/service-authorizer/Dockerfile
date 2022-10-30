FROM golang:1.19 AS build-image

WORKDIR /app

COPY . .

RUN env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/authorizer *.go

FROM amazon/aws-lambda-go:latest

COPY --from=build-image /app/bin ${LAMBDA_TASK_ROOT}/bin/

CMD [ "app" ]