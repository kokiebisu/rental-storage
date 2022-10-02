#! /bin/sh

# Go
(cd packages/service-authorizer && make prepare);
# Typescript
turbo run build