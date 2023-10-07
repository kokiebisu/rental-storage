resource "aws_ecr_repository" "service_search" {
  force_delete = true
  name         = "ecs-service-search-${var.environment}"
}
