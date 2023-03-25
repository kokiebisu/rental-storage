NAMESPACE = @server-mock

.PHONY: deps-start
deps-start:
	@echo "[ $(NAMESPACE) ] Running Mock Server..."
	@docker-compose up -d --build

.PHONY: deps-stop
deps-stop:
	@echo "[ $(NAMESPACE) ] Stopping Mock Server..."
	@docker-compose down --volumes