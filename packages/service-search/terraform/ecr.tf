data "aws_ecr_repository" "service_search" {
  name = "ecs-service-search-${var.environment}"
}

resource "null_resource" "docker_push" {
  triggers = {
    timestamp = timestamp()
  }

  provisioner "local-exec" {
    command = <<EOT
      # Build the Docker image
      docker build -t app:latest ../

      # Log in to the ECR registry
      aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin ${data.aws_ecr_repository.service_search.repository_url}

      # Tag the Docker image
      docker tag app:latest ${data.aws_ecr_repository.service_search.repository_url}:latest

      # Push the Docker image to the ECR repository
      docker push ${data.aws_ecr_repository.service_search.repository_url}:latest
    EOT
  }
}