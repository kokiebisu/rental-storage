install-deps-local:
	make start-venv && pip install -t vendor -r requirements.txt

start-venv:
	source venv/bin/activate

.PHONY: test
test:
	pytest -v test/unit_test

.PHONY: lint
lint:
	flake8 . --exclude ./vendor/**,./node_modules/**
