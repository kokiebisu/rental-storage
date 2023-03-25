NAMESPACE = @service-slack

.PHONY: set-profile
set-profile:
	@echo "[ $(NAMESPACE) ] Setting AWS_PROFILE to rental-storage..."
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
	@pytest -v test/unit_test

.PHONY: lint
lint:
	@echo "[ $(NAMESPACE) ] Linting..."
	@flake8 . --exclude ./vendor/**,./node_modules/**
