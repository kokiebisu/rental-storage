echo "Removing service-authentication bin folder..."
(cd packages/service-authentication && make clean);
echo "Removing service-authorizer bin folder..."
(cd packages/service-authorizer && make clean);
echo "Removing service-booking dep/dist/turbo folders..."
(cd packages/service-booking && rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo); 
echo "Removing service-image dep/dist/turbo folders..."
(cd packages/service-image && rm -rf node_modules); 
echo "Removing service-listing dep/dist/turbo folders..."
(cd packages/service-listing && rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo); 
echo "Removing service-user dep/dist/turbo folders..."
(cd packages/service-user && rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo);
echo "Removing service-item dep/dist/turbo folders..."
(cd packages/service-item && rm -rf node_modules && rm -rf dist && rm -rf .turbo && rm tsconfig.tsbuildinfo);