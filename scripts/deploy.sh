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
echo "Install dev dependencies for appsync" && 
# cd api/appsync && npm install -D && cd ../.. &&
cd infrastructure && serverless deploy --config serverless.dev.yml --stage dev
# echo "Install dev dependencies for apigateway" && 
# cd api/apigateway && npm install -D && cd ../.. &&

# echo "Install dev dependencies for infrastructure" && 
# cd infrastructure && npm install -D && cd .. &&

# echo "Install dev depdendencies for authentication service" &&
# cd packages/service-authentication && npm install -D && cd ../.. &&
# echo "Install dev depdendencies for authorizer service" &&
# cd packages/service-authorizer && npm install -D && cd ../.. &&
# echo "Install dev depdendencies for booking service" &&
# cd packages/service-booking && npm install -D && cd ../.. &&
# echo "Install dev dependencies for image service" &&
# cd packages/service-image && npm install -D && cd ../.. &&
# echo "Install dev dependencies for listing service" &&
# cd packages/service-listing && npm install -D && cd ../.. &&
# echo "Install dev dependencies for payment service" &&
# cd packages/service-payment && npm install -D && cd ../.. &&
# echo "Install dev dependencies for slack service" &&
# cd packages/service-slack && npm install -D && cd ../.. &&
# echo "Install dev dependencies for user service" &&
# cd packages/service-user && npm install -D && cd ../.. &&

# echo "Listing the files..."
# ls

# TEST=$(pwd)
# echo $TEST
# echo "Deploy $ENVIRONMENT" && 
# serverless deploy --stage $ENVIRONMENT;