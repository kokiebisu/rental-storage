install-deps:
	./scripts/install.sh

.PHONY: build
build: 
	npx turbo run build && 
	cd packages/service-payment && make build && cd ../.. &&
	cd packages/service-authentication && make build && cd ../.. &&
	cd packages/service-authorizer && make build && cd ../..

.PHONY: enable-precommit
enable-precommit:
	cp scripts/hooks/pre-commit.sh .git/hooks/pre-commit

.PHONY: disable-precommit
disable-precommit:
	rm .git/hooks/pre-commit
