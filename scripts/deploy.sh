#! /bin/sh

ENVIRONMENT=$1
echo "ENVIRONMENT: $ENVIRONMENT"

if [ -e serverless-compose.yml ]; then
    echo "Removing serverless-compose.yml";
    rm serverless-compose.yml ;
fi

if [ "$ENVIRONMENT" = "dev" ]; then
    echo "Copying serverless.dev.yml";
    cp config/serverless-compose.dev.yml ./serverless-compose.yml;
else
    echo "Copying serverless.yml";
    cp config/serverless-compose.yml ./serverless-compose.yml;
fi

# Install dev dependencies
cd api/appsync && npm ci --only=dev && cd ../.. &&
cd api/apigateway && npm ci --only=dev && cd ../.. &&

cd infrastructure && npm ci --only=dev && cd ../.. &&

cd packages/service-authentication && npm ci --only=dev && cd ../.. &&
cd packages/service-authorizer && npm ci --only=dev && cd ../.. &&
cd packages/service-booking && npm ci --only=dev && cd ../.. &&
cd packages/service-image && npm ci --only=dev && cd ../.. &&
cd packages/service-listing && npm ci --only=dev && cd ../.. &&
cd packages/service-payment && npm ci --only=dev && cd ../.. &&
cd packages/service-slack && npm ci --only=dev && cd ../.. &&
cd packages/service-user && npm ci --only=dev && cd ../.. &&

npx sls deploy --stage $ENVIRONMENT;