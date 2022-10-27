#! /bin/sh

ENVIRONMENT=$1
echo "ENVIRONMENT: $ENVIRONMENT"

# if [ -e serverless-compose.yml ]; then
#     echo "Removing serverless-compose.yml";
#     rm serverless-compose.yml ;
# fi

# if [ "$ENVIRONMENT" = "dev" ]; then
#     echo "Copying serverless.dev.yml";
#     cp config/serverless-compose.dev.yml ./serverless-compose.yml;
# else
#     echo "Copying serverless.yml";
#     cp config/serverless-compose.yml ./serverless-compose.yml;
# fi

# Install dev dependencies
echo "Deploy infrastructure" && 
# cd api/appsync && npm install -D && cd ../.. &&
cd infrastructure && serverless deploy --config serverless.dev.yml --stage dev && cd .. &&
echo "Deploy apigateway" && 
cd api/apigateway && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&

echo "Deploy authentication service" &&
cd packages/service-authentication && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&

echo "Deploy authorizer service" &&
cd packages/service-authorizer && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&
echo "Deploy booking service" &&
cd packages/service-booking && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&
echo "Deploy image service" &&
cd packages/service-image && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&
echo "Deploy listing service" &&
cd packages/service-listing && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&
echo "Deploy payment service" &&
cd packages/service-payment && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&
echo "Deploy slack service" &&
cd packages/service-slack && serverless deploy --config serverless.dev.yml --stage dev && cd ../.. &&
echo "Deploy user service" &&
cd packages/service-user && serverless deploy --config serverless.dev.yml --stage dev