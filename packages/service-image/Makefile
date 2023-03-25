.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage"
	@export AWS_PROFILE=rental-storage

install-deps-local:
	@echo "Installing dependencies locally"
	@make start-venv && pip install -t vendor -r requirements.txt

start-venv:
	@echo "Starting virtual environment"
	@source venv/bin/activate

.PHONY: test-unit
test-unit:
	@echo "Running unit tests"
	@python -m pytest -v test/unit_test

.PHONY: integration-test
integration-test:
	@echo "Running integration tests"
	@python -m pytest -v test/integration_test
