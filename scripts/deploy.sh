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
cd api/appsync && npm install -D && cd ../.. &&
cd api/apigateway && npm install -D && cd ../.. &&

cd infrastructure && npm install -D && cd ../.. &&

cd packages/service-authentication && npm install -D && cd ../.. &&
cd packages/service-authorizer && npm install -D && cd ../.. &&
cd packages/service-booking && npm install -D && cd ../.. &&
cd packages/service-image && npm install -D && cd ../.. &&
cd packages/service-listing && npm install -D && cd ../.. &&
cd packages/service-payment && npm install -D && cd ../.. &&
cd packages/service-slack && npm install -D && cd ../.. &&
cd packages/service-user && npm install -D && cd ../.. &&

npx sls deploy --stage $ENVIRONMENT;