#! /bin/sh

ENVIRONMENT=$1
echo "ENVIRONMENT: $ENVIRONMENT"

if [ -e serverless-compose.yml ]; then
    echo "Removing serverless-compose.yml"
    rm serverless-compose.yml 
fi

if [ "$ENVIRONMENT" = "dev" ]; then
    echo "Copying serverless.dev.yml"
    cp config/serverless-compose.dev.yml ./serverless-compose.yml;
else
    echo "Copying serverless.yml"
    cp config/serverless-compose.yml ./serverless-compose.yml;
fi

cd api/appsync && npm install && cd ../..;

npx sls deploy --stage $ENVIRONMENT;