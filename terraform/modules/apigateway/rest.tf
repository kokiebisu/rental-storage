resource "aws_api_gateway_rest_api" "rest" {
    name = "rest-api"

    tags = {
      "Name" = "rest-api"
    }
}

resource "aws_api_gateway_resource" "health" {
  rest_api_id = aws_api_gateway_rest_api.rest.id
  parent_id   = aws_api_gateway_rest_api.rest.root_resource_id
  path_part   = "health"
}

resource "aws_api_gateway_method" "health_get" {
  rest_api_id   = aws_api_gateway_rest_api.rest.id
  resource_id   = aws_api_gateway_resource.health.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "health" {
  depends_on = [
    aws_api_gateway_rest_api.rest,
    aws_api_gateway_resource.health,
    aws_api_gateway_method.health_get
  ]
  http_method = aws_api_gateway_method.health_get.http_method
  resource_id = aws_api_gateway_resource.health.id
  rest_api_id = aws_api_gateway_rest_api.rest.id
  type        = "MOCK"
}

resource "aws_api_gateway_stage" "this" {
  deployment_id = aws_api_gateway_deployment.this.id
  rest_api_id   = aws_api_gateway_rest_api.rest.id
  stage_name    = "${var.environment}"
}

resource "aws_api_gateway_deployment" "this" {
  depends_on = [aws_api_gateway_resource.health, aws_api_gateway_integration.health, aws_api_gateway_method.health_get]
  rest_api_id = aws_api_gateway_rest_api.rest.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.rest.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}
