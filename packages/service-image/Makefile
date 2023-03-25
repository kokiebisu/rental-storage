NAMESPACE = @service-image

.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage..."
	@export AWS_PROFILE=rental-storage

install-deps-local:
	@echo "[ $(NAMESPACE) ] Installing dependencies locally..."
	@make start-venv && pip install -t vendor -r requirements.txt

start-venv:
	@echo "[ $(NAMESPACE) ] Starting virtual environment..."
	@source venv/bin/activate

.PHONY: test-unit
test-unit:
	@echo "[ $(NAMESPACE) ] Running unit tests..."
	@python -m pytest -v test/unit_test

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running integration tests..."
	@python -m pytest -v test/integration_test
