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

.PHONY: setup-terraform
setup-terraform:
	@echo "[ $(NAMESPACE) ] Setting up terraform..."
	@./scripts/setup-terraform.sh
