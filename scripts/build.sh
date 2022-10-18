#! /bin/sh

# Check if docker is running
if ! docker info > /dev/null 2>&1; then
  echo "This script uses docker, and it isn't running - please start docker and try again!"
  exit 1
fi

# Go
(cd packages/service-authorizer && make prepare || { echo 'failed' ; exit 1; });
(cd packages/service-authentication && make prepare || { echo 'failed' ; exit 1; });
(cd packages/service-payment && make prepare || { echo 'failed' ; exit 1; });

# Typescript
(turbo run build || { echo 'failed' ; exit 1; });