resource "aws_ssm_parameter" "rest_api_id" {
    name = "/terraform/apigateway/rest-api-id"
    type = "String"
    value = aws_api_gateway_rest_api.default.id
}

resource "aws_ssm_parameter" "rest_api_root_resource_id" {
    name = "/terraform/apigateway/rest-api-root-resource-id"
    type = "String"
    value = aws_api_gateway_rest_api.default.root_resource_id
}