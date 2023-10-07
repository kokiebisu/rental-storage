resource "aws_ecs_cluster" "default" {
  name = "${var.environment}-app"
}
