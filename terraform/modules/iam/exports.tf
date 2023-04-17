resource "aws_ssm_parameter" "lambda_role_arn" {
    name = "/terraform/iam/lambda-role-arn"
    type = "String"
    value = aws_iam_role.lambda.arn
}

resource "aws_ssm_parameter" "service_chat_lambda_role_arn" {
    name = "/terraform/iam/service-chat-lambda-role-arn"
    type = "String"
    value = aws_iam_role.service_chat.name
}
