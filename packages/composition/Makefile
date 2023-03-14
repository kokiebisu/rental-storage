.PHONY: lint
lint: 
	npx eslint .

.PHONY: test-unit
test-unit: 
	npx jest --config test/unit/jest.config.ts --coverage

.PHONY: test-api
test-api: 
	npx jest --config test/api/jest.config.ts --coverage

.PHONY: remove
remove:
	npx -y serverless@3.28.1 remove --config serverless.yml --stage local

.PHONY: deploy
deploy:
	npx -y serverless@3.28.1 deploy --config serverless.yml --stage local
