install-deps-local:
	make start-venv && pip install -t vendor -r requirements.txt

start-venv:
	source venv/bin/activate

.PHONY: test
test:
	python -m pytest -v