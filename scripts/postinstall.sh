# Go
(cd packages/service-authorizer && make prepare || { echo 'failed' ; exit 1; });
(cd packages/service-authentication && make prepare || { echo 'failed' ; exit 1; });
(cd packages/service-payment && make prepare || { echo 'failed' ; exit 1; });

# Typescript
(cd packages/service-booking && npm install || { echo 'failed' ; exit 1; }); 
(cd packages/service-image && npm install || { echo 'failed' ; exit 1; }); 
(cd packages/service-listing && npm install || { echo 'failed' ; exit 1; }); 
(cd packages/service-user && npm install || { echo 'failed' ; exit 1; });

(cd infrastructure && npm install || { echo 'failed' ; exit 1; });

(cd api/appsync && npm install || { echo 'failed' ; exit 1; });
(cd api/apigateway && npm install || { echo 'failed' ; exit 1; });