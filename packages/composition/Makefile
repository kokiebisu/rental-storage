.PHONY: lint
lint: 
	npx eslint .

.PHONY: test-unit
test-unit: 
	npx jest --config test/unit/jest.config.ts --coverage

.PHONY: test-api
test-api: 
	npx jest --config test/api/jest.config.ts --coverage
