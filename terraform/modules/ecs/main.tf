# resource "aws_ecs_cluster" "search" {
#     name = "service-search"
# }

# resource "aws_ecs_task_definition" "search" {
#     family = "search"
#     container_definitions = <<EOF
#     [
#         {
#             "name": "container",
#             "image": "", // ECR Registry
#             "portMappings": [
#                 {
#                     "containerPort": 8080,
#                     "hostPort": 0
#                 }
#             ]
#         }
#     ]
#     EOF
# }

## need a task execution role to access opensearch

## need permission to pull from ecr registry
