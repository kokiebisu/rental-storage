install-deps-local:
	make start-venv && pip install -t vendor -r requirements.txt

start-venv:
	source venv/bin/activate

.PHONY: test-unit
test-unit:
	pytest -v test/unit_test

.PHONY: lint
lint:
	flake8 . --exclude ./vendor/**,./node_modules/**

.PHONY: remove
remove:
	npx -y serverless@3.28.1 remove --config serverless.yml --stage local

.PHONY: deploy
deploy:
	npx -y serverless@3.28.1 deploy --config serverless.yml --stage local
