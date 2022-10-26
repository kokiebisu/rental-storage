#! /bin/sh
echo "Installing package root dependencies..."
npm install

echo "Installing Infrastructure dependencies..."
cd infrastructure && npm ci --only=prod && cd ..

# api
echo "Installing Appsync dependencies..."
cd api/appsync && npm ci --only=prod && cd ../..
echo "Installing Api Gateway dependencies..."
cd api/apigateway && npm ci --only=prod && cd ../..


# typescript
echo "Installing Booking Service dependencies..."
cd packages/service-booking && npm ci --only=prod && cd ../..
echo "Installing Listing Service dependencies..."
cd packages/service-listing && npm ci --only=prod && cd ../..
echo "Installing User Service dependencies..."
cd packages/service-user && npm ci --only=prod && cd ../..

# python
echo "Installing Image Service dependencies..."
cd packages/service-image && npm ci --only=prod && cd ../..
echo "Installing Slack Service dependencies..."
cd packages/service-slack && npm ci --only=prod && cd ../..

# golang
echo "Installing Authentication Service dependencies..."
cd packages/service-authentication && npm ci --only=prod && cd ../..
echo "Installing Authorizer Service dependencies..."
cd packages/service-authorizer && npm ci --only=prod && cd ../..
echo "Installing Payment Service dependencies..."
cd packages/service-payment && npm ci --only=prod && cd ../..


