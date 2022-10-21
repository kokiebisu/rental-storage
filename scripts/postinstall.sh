#! /bin/sh

# Go
# (cd packages/service-authorizer && make prepare || { echo 'failed' ; exit 1; });
# (cd packages/service-authentication && make prepare || { echo 'failed' ; exit 1; });
# (cd packages/service-payment && make prepare || { echo 'failed' ; exit 1; });

# Typescript
echo "Setting up Booking Service..."
cd packages/service-booking && npm ci && cd ../..
echo "Setting up Image Service..."
cd packages/service-image && npm ci && cd ../..
echo "Setting up Listing Service..."
cd packages/service-listing && npm ci && cd ../..
echo "Setting up User Service..."
cd packages/service-user && npm ci && cd ../..

echo "Setting up Infrastructure..."
cd infrastructure && npm ci && cd ../..

echo "Setting up Appsync..."
cd api/appsync && npm ci && cd ../..
echo "Setting up Api Gateway..."
cd api/apigateway && npm ci && cd ../..