#! /bin/sh

# Go
(cd packages/service-authorizer && make prepare || { echo 'failed' ; exit 1; });
(cd packages/service-authentication && make prepare || { echo 'failed' ; exit 1; });

# Typescript
(turbo run build || { echo 'failed' ; exit 1; });