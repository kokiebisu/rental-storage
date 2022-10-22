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
echo "Setting up Authorizer Service..."
cd packages/service-authorizer && npm ci && cd ../..
echo "Setting up Authentication Service..."
cd packages/service-authentication && npm ci && cd ../..
echo "Setting up Payment Service..."
cd packages/service-payment && npm ci && cd ../..