FROM openjdk:8-jdk-alpine

WORKDIR /app

COPY target/search-service-app-*.jar /app/app.jar
COPY config.yml .

EXPOSE 8080

CMD ["java", "-jar", "app.jar", "server", "config.yml"]