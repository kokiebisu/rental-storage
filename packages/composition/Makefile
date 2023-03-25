.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage"
	@export AWS_PROFILE=rental-storage

.PHONY: lint
lint:
	@echo "Linting"
	@npx eslint .

.PHONY: test-unit
test-unit: 
	@echo "Running unit tests"
	@npx jest --config test/unit/jest.config.ts --coverage

.PHONY: test-api
test-api: 
	@echo "Running API tests"
	@npx jest --config test/api/jest.config.ts --coverage
