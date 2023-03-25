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
