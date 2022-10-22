#! /bin/sh
echo "Setting up package root..."
npm ci

echo "Setting up Booking Service..."
cd packages/service-booking && npm ci && cd ../..
echo "Setting up Image Service..."
cd packages/service-image && npm ci && cd ../..
echo "Setting up Listing Service..."
cd packages/service-listing && npm ci && cd ../..
echo "Setting up User Service..."
cd packages/service-user && npm ci && cd ../..

echo "Setting up Infrastructure..."
cd infrastructure && npm ci && cd ..

echo "Setting up Appsync..."
cd api/appsync && npm ci && cd ../..
echo "Setting up Api Gateway..."
cd api/apigateway && npm ci && cd ../..

# Go
cd packages/service-authorizer && npm ci && cd ../..
# (cd packages/service-authentication && make prepare || { echo 'failed' ; exit 1; });
# (cd packages/service-payment && make prepare || { echo 'failed' ; exit 1; });
