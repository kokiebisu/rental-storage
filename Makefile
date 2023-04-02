NAMESPACE = @root

.PHONY: enable-precommit
enable-precommit:
	@echo "[ $(NAMESPACE) ] Enabling pre-commit hook..."
	@cp scripts/hooks/pre-commit.sh .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit

.PHONY: disable-precommit
disable-precommit:
	@echo "[ $(NAMESPACE) ] Disabling pre-commit hook..."
	@rm .git/hooks/pre-commit

.PHONY: destroy
destroy:
	@echo "[ $(NAMESPACE) ] Destroying project..."
	@./scripts/destroy.sh

.PHONY: destroy-service
destroy-service:
	@echo "[ $(NAMESPACE) ] Destroying service..."
	@./scripts/destroy-service.sh

.PHONY: destroy-terraform
destroy-terraform:
	@echo "Checking for environment..."
	@if [ -z "$(ENV)" ]; then \
		echo "Please pass the environment as an argument."; \
		echo "e.g. make destroy-terraform ENV=local"; \
		exit 1; \
	fi
	@echo "[ $(NAMESPACE) ] Destroying terraform..."
	@./scripts/destroy-terraform.sh $(ENV)

.PHONY: setup
setup:
	@echo "[ $(NAMESPACE) ] Setting up project..."
	@./scripts/setup.sh

.PHONY: setup-deps
setup-deps:
	@echo "[ $(NAMESPACE) ] Setting up dependencies..."
	@./scripts/setup-deps.sh

.PHONY: setup-service
setup-service:
	@echo "[ $(NAMESPACE) ] Setting up service..."
	@./scripts/setup-service.sh

# pass the environment as an argument
# e.g. make setup-terraform ENV=local
.PHONY: setup-terraform
setup-terraform:
	@echo "Checking for environment..."
	@if [ -z "$(ENV)" ]; then \
		echo "Please pass the environment as an argument."; \
		echo "e.g. make setup-terraform ENV=local"; \
		exit 1; \
	fi
	@echo "[ $(NAMESPACE) ] Setting up terraform..."
	@./scripts/setup-terraform.sh $(ENV)

# pass the environment as an argument
# e.g. make setup-terraform-config ENV=local
.PHONY: setup-terraform-config
setup-terraform-config:
	@echo "Checking for environment..."
	@if [ -z "$(ENV)" ]; then \
		echo "Please pass the environment as an argument."; \
		echo "e.g. make setup-terraform-config ENV=local"; \
		exit 1; \
	fi
	@echo "[ $(NAMESPACE) ] Setting up terraform config..."
	@./scripts/setup-terraform-config.sh $(ENV)
