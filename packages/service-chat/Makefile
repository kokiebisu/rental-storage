NAMESPACE = @service-chat

.PHONY: set-profile
set-profile:
	@echo "Setting AWS_PROFILE to rental-storage..."
	@export AWS_PROFILE=rental-storage

.PHONY: install-deps-local
install-deps-local:
	@echo "[ $(NAMESPACE) ] Installing dependencies locally..."
	@pip install -t vendor -r requirements.txt

.PHONY: test-unit
test-unit:
	@echo "[ $(NAMESPACE) ] Running unit tests..."
	@python -m pytest -v test/unit_test

.PHONY: test-integration
test-integration:
	@echo "[ $(NAMESPACE) ] Running integration tests..."
	@python -m pytest -v test/integration_test
