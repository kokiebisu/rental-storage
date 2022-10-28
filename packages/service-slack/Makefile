install-deps-local:
	make start-venv && pip install -t vendor -r requirements.txt

start-venv:
	source venv/bin/activate

.PHONY: unit-test
unit-test:
	python -m pytest -v test/unit_test

.PHONY: integration-test
integration-test:
	python -m pytest -v test/integration_test