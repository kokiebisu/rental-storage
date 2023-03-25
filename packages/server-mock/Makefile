NAMESPACE = @server-mock

.PHONY: deps-start
deps-start:
	@echo "[ $(NAMESPACE) ] Running mock server..."
	@docker-compose up -d --build

.PHONY: deps-stop
deps-stop:
	@echo "[ $(NAMESPACE) ] Stopping mock server..."
	@docker-compose down --volumes --remove-orphans