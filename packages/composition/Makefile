.PHONY: lint
lint: 
	npx eslint .

.PHONY: test-unit
test-unit: 
	npx jest --config test/unit/jest.config.ts --coverage

.PHONY: test-api
test-api: 
	npx jest --config test/api/jest.config.ts --coverage

.PHONY: remote-remove
remote-remove:
	npx serverless@3.28.1 remove --config serverless.yml --stage local

.PHONY: remote-deploy
remote-deploy:
	npx serverless@3.28.1 deploy --config serverless.yml --stage local
