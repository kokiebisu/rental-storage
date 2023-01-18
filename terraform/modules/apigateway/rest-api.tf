resource "aws_api_gateway_rest_api" "default" {
    name = "ApiGateway"

    tags = {
      "Name" = "ApiGateway"
    }
}

resource "aws_api_gateway_resource" "health" {
  rest_api_id = aws_api_gateway_rest_api.default.id
  parent_id   = aws_api_gateway_rest_api.default.root_resource_id
  path_part   = "health"
}

resource "aws_api_gateway_method" "health_get" {
  rest_api_id   = aws_api_gateway_rest_api.default.id
  resource_id   = aws_api_gateway_resource.health.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "health" {
  depends_on = [
    aws_api_gateway_rest_api.default,
    aws_api_gateway_resource.health,
    aws_api_gateway_method.health_get
  ]
  http_method = aws_api_gateway_method.health_get.http_method
  resource_id = aws_api_gateway_resource.health.id
  rest_api_id = aws_api_gateway_rest_api.default.id
  type        = "MOCK"
}

resource "aws_api_gateway_stage" "default" {
  deployment_id = aws_api_gateway_deployment.default.id
  rest_api_id   = aws_api_gateway_rest_api.default.id
  stage_name    = "${var.environment}"
}

resource "aws_api_gateway_deployment" "default" {
  depends_on = [aws_api_gateway_resource.health, aws_api_gateway_integration.health, aws_api_gateway_method.health_get]
  rest_api_id = aws_api_gateway_rest_api.default.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.default.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}
