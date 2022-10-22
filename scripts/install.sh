#! /bin/sh
echo "Installing package root dependencies..."
npm ci

# typescript
echo "Installing Booking Service dependencies..."
cd packages/service-booking && npm ci && cd ../..
echo "Installing Listing Service dependencies..."
cd packages/service-listing && npm ci && cd ../..
echo "Installing User Service dependencies..."
cd packages/service-user && npm ci && cd ../..

# python
echo "Installing Image Service dependencies..."
cd packages/service-image && npm ci && cd ../..
echo "Installing Slack Service dependencies..."
cd packages/service-slack && npm ci && cd ../..

# golang
echo "Installing Authentication Service dependencies..."
cd packages/service-authentication && npm ci && cd ../..
echo "Installing Authorizer Service dependencies..."
cd packages/service-authorizer && npm ci && cd ../..
echo "Installing Payment Service dependencies..."
cd packages/service-payment && npm ci && cd ../..

echo "Installing Infrastructure dependencies..."
cd infrastructure && npm ci && cd ..

echo "Installing Appsync dependencies..."
cd api/appsync && npm ci && cd ../..
echo "Installing Api Gateway dependencies..."
cd api/apigateway && npm ci && cd ../..

