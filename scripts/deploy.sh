#! /bin/sh

ENVIRONMENT=$1
echo "ENVIRONMENT: $ENVIRONMENT"

if [ test -f "serverless.yml" ]; then
    echo "Removing serverless.yml"
    rm serverless.yml 
fi

if [ "$ENVIRONMENT" == "dev" ]; then
    echo "Copying serverless.dev.yml"
    cp config/serverless-compose.dev.yml ./serverless-compose.yml;
else
    echo "Copying serverless.yml"
    cp config/serverless-compose.yml ./serverless-compose.yml;
fi
npx sls deploy --stage $ENVIRONMENT;