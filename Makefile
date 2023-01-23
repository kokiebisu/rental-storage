.PHONY: enable-precommit
enable-precommit:
	cp scripts/hooks/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

.PHONY: disable-precommit
disable-precommit:
	rm .git/hooks/pre-commit
