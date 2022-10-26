#! /bin/sh
echo "Installing package root dependencies..."
npm install

echo "Installing Infrastructure dependencies..."
cd infrastructure && npm install --omit=dev && cd ..

# api
echo "Installing Appsync dependencies..."
cd api/appsync && npm install --omit=dev && cd ../..
echo "Installing Api Gateway dependencies..."
cd api/apigateway && npm install --omit=dev && cd ../..


# typescript
echo "Installing Booking Service dependencies..."
cd packages/service-booking && npm install --omit=dev && cd ../..
echo "Installing Listing Service dependencies..."
cd packages/service-listing && npm install --omit=dev && cd ../..
echo "Installing User Service dependencies..."
cd packages/service-user && npm install --omit=dev && cd ../..

# python
echo "Installing Image Service dependencies..."
cd packages/service-image && npm install --omit=dev && cd ../..
echo "Installing Slack Service dependencies..."
cd packages/service-slack && npm install --omit=dev && cd ../..

# golang
echo "Installing Authentication Service dependencies..."
cd packages/service-authentication && npm install --omit=dev && cd ../..
echo "Installing Authorizer Service dependencies..."
cd packages/service-authorizer && npm install --omit=dev && cd ../..
echo "Installing Payment Service dependencies..."
cd packages/service-payment && npm install --omit=dev && cd ../..


