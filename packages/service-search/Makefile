# Variables
JAR_FILE=target/service-search-1.0.0.jar

# Build the project
build:
	mvn clean package

# Run the server
run:
	java -jar $(JAR_FILE) server config.yml

# Stop the server
stop:
	pkill -f $(JAR_FILE)

# Clean build artifacts
clean:
	mvn clean
