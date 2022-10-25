#! /bin/sh
echo "Installing package root dependencies..."
npm install

# typescript
echo "Installing Booking Service dependencies..."
cd packages/service-booking && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..
echo "Installing Listing Service dependencies..."
cd packages/service-listing && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..
echo "Installing User Service dependencies..."
cd packages/service-user && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..

# python
echo "Installing Image Service dependencies..."
cd packages/service-image && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..
echo "Installing Slack Service dependencies..."
cd packages/service-slack && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..

# golang
echo "Installing Authentication Service dependencies..."
cd packages/service-authentication && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..
echo "Installing Authorizer Service dependencies..."
cd packages/service-authorizer && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..
echo "Installing Payment Service dependencies..."
cd packages/service-payment && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..

echo "Installing Infrastructure dependencies..."
cd infrastructure && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ..

echo "Installing Appsync dependencies..."
cd api/appsync && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..
echo "Installing Api Gateway dependencies..."
cd api/apigateway && npm install && npx serverless plugin install -n serverless-deployment-bucket serverless-export-env && cd ../..

