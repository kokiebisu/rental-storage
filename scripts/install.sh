#! /bin/sh
echo "Installing package root dependencies..."
npm install

echo "Installing Infrastructure dependencies..."
cd infrastructure && npm install && cd ..

echo "Installing Appsync dependencies..."
cd api/appsync && npm install && cd ../..
echo "Installing Api Gateway dependencies..."
cd api/apigateway && npm install && cd ../..


# typescript
echo "Installing Booking Service dependencies..."
cd packages/service-booking && npm install && cd ../..
echo "Installing Listing Service dependencies..."
cd packages/service-listing && npm install && cd ../..
echo "Installing User Service dependencies..."
cd packages/service-user && npm install && cd ../..

# python
echo "Installing Image Service dependencies..."
cd packages/service-image && npm install && cd ../..
echo "Installing Slack Service dependencies..."
cd packages/service-slack && npm install && cd ../..

# golang
echo "Installing Authentication Service dependencies..."
cd packages/service-authentication && npm install && cd ../..
echo "Installing Authorizer Service dependencies..."
cd packages/service-authorizer && npm install && cd ../..
echo "Installing Payment Service dependencies..."
cd packages/service-payment && npm install && cd ../..


