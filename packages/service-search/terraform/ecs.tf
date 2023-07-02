resource "aws_ecs_cluster" "search" {
    name = "service-search"
}

resource "aws_ecs_service" "search" {
  name            = "service"
  cluster         = aws_ecs_cluster.search.id
  task_definition = aws_ecs_task_definition.search.arn
  desired_count   = 1
  launch_type = "FARGATE"

  network_configuration {
    security_groups = [data.aws_security_group.ecs_service.id]
    subnets         = [data.aws_subnet.c.id, data.aws_subnet.d.id]
  }
}

resource "aws_ecs_task_definition" "search" {
    family = "search"
    network_mode = "awsvpc"
    requires_compatibilities = ["FARGATE"]
    cpu                      = "256"
    memory = "512"

    execution_role_arn       = data.aws_iam_role.ecs_task_execution_role.arn

    container_definitions = <<EOF
    [
        {
            "name": "java-container",
            "image": "${data.aws_ecr_repository.service_search.repository_url}:latest",
            "cpu": 256,
            "memory": 512,
            "essential": true,
            "portMappings": [
                {
                "containerPort": 8080,
                "hostPort": 8080
                }
            ]
        }
    ]
    EOF
}
