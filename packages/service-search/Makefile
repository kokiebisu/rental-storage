# Variables
JAR_FILE=target/search-service-app-1.0.0.jar

# Build the project
build:
	mvn clean && mvn clean package

# Run the server
run:
	java -jar $(JAR_FILE) server config.yml

# Stop the server
stop:
	pkill -f $(JAR_FILE)

# Clean build artifacts
clean:
	mvn clean

# pass the environment as an argument
# e.g. make setup-terraform ENV=local
.PHONY: setup-terraform
setup-terraform:
	@echo "Checking for environment..."
	@if [ -z "$(ENV)" ]; then \
		echo "Please pass the environment as an argument."; \
		echo "e.g. make setup-terraform ENV=local"; \
		exit 1; \
	fi
	@echo "[ $(NAMESPACE) ] Setting up terraform..."
	@cd terraform && terraform apply -auto-approve -var-file=terraform.tfvars

.PHONY: destroy-terraform
destroy-terraform:
	@echo "Checking for environment..."
	@if [ -z "$(ENV)" ]; then \
		echo "Please pass the environment as an argument."; \
		echo "e.g. make destroy-terraform ENV=local"; \
		exit 1; \
	fi
	@echo "[ $(NAMESPACE) ] Destroying terraform..."
	@cd terraform && terraform destroy -auto-approve -var-file=terraform.tfvars