install-deps:
	./scripts/install.sh

.PHONY: build
build: 
	npx turbo run build && 
	cd packages/service-payment && make build && cd ../.. &&
	cd packages/service-authentication && make build && cd ../.. &&
	cd packages/service-authorizer && make build && cd ../..

.PHONY: deploy-local
deploy-local:
	./scripts/deploy.sh local

.PHONY: deploy-dev
deploy-dev:
	./scripts/deploy.sh dev