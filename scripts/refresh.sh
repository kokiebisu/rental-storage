echo "Removing service-authentication bin folder..."
(cd packages/service-authentication && make clean);
echo "Removing service-authorizer bin folder..."
(cd packages/service-authorizer && make clean);
echo "Removing service-payment bin folder..."
(cd packages/service-payment && make clean);

echo "Removing service-booking dep/dist/turbo folders..."
(cd packages/service-booking && make clean); 
echo "Removing service-image dep/dist/turbo folders..."
(cd packages/service-image && make clean); 
echo "Removing service-listing dep/dist/turbo folders..."
(cd packages/service-listing && make clean); 
echo "Removing service-user dep/dist/turbo folders..."
(cd packages/service-user && make clean);

echo "Removing service-user dep/dist/turbo folders..."
(cd packages/service-user && make clean);

echo "Removing infrastructure dep/dist/turbo folders..."
(cd packages/service-user && make clean);

echo "Removing api-apigateway dep folder..."
(cd api/apigateway && make clean);
echo "Removing api-appsync dep folder..."
(cd api/appsync && make clean);
