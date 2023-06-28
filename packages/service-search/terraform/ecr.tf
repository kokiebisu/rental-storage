resource "aws_ecr_repository" "service_search" {
  name = "ecs-service-search-${var.environment}"
}
