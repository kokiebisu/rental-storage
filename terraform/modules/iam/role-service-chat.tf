resource "aws_iam_role_policy_attachment" "enable_service_chat_log" {
  role       = aws_iam_role.service_chat.name
  policy_arn = aws_iam_policy.enable_log.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_chat_api_gateway" {
  role       = aws_iam_role.service_chat.name
  policy_arn = aws_iam_policy.enable_api_gateway_invoke.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_chat_dynamodb" {
  role       = aws_iam_role.service_chat.name
  policy_arn = aws_iam_policy.enable_dynamodb.arn
}

resource "aws_iam_role" "service_chat" {
  name                  = "service-chat-lambda-role"
  force_detach_policies = true
  assume_role_policy = file("${local.path}/role-lambda.tpl")
}
