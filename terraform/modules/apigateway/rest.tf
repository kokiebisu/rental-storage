resource "aws_api_gateway_rest_api" "rest" {
  name = "rest-api"

  tags = {
    "Name" = "rest-api"
  }
}

##### Health
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
  rest_api_id = aws_api_gateway_rest_api.rest.id
  resource_id = aws_api_gateway_resource.health.id
  http_method = aws_api_gateway_method.health_get.http_method
  type        = "MOCK"
}

##### Search Service
resource "aws_api_gateway_resource" "service_search" {
  rest_api_id = aws_api_gateway_rest_api.rest.id
  parent_id   = aws_api_gateway_rest_api.rest.root_resource_id
  path_part   = "search"
}

resource "aws_api_gateway_method" "service_search" {
  rest_api_id   = aws_api_gateway_rest_api.rest.id
  resource_id   = aws_api_gateway_resource.service_search.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "service_search" {
  depends_on = [
    aws_api_gateway_rest_api.rest,
    aws_api_gateway_resource.service_search,
    aws_api_gateway_method.service_search
  ]
  rest_api_id             = aws_api_gateway_rest_api.rest.id
  resource_id             = aws_api_gateway_resource.service_search.id
  http_method             = aws_api_gateway_method.service_search.http_method
  type                    = "HTTP_PROXY"
  uri                     = "http://${var.lb_dns_name}"
  integration_http_method = "ANY"
}

resource "aws_api_gateway_stage" "this" {
  deployment_id = aws_api_gateway_deployment.this.id
  rest_api_id   = aws_api_gateway_rest_api.rest.id
  stage_name    = var.environment
}

resource "aws_api_gateway_deployment" "this" {
  depends_on = [
    aws_api_gateway_resource.health,
    aws_api_gateway_integration.health,
    aws_api_gateway_method.health_get,
    aws_api_gateway_resource.service_search,
    aws_api_gateway_integration.service_search,
    aws_api_gateway_method.service_search
  ]
  rest_api_id = aws_api_gateway_rest_api.rest.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.rest.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}
