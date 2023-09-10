data "aws_iam_role" "ecs_task_execution_role" {
    name = "ecs-task-execution-role"
}

data "aws_subnet" "c" {
    availability_zone = "us-east-1c"

    filter {
        name   = "tag:Name"
        values = ["subnet-c"]
    }
}

data "aws_subnet" "d" {
    availability_zone = "us-east-1d"

    filter {
        name   = "tag:Name"
        values = ["subnet-d"]
    }
}

data "aws_security_group" "ecs_service" {
    filter {
      name = "tag:Name"
      values = ["ecs"]
    }
}
