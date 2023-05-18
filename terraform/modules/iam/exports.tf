resource "aws_ssm_parameter" "terraform_user_access_key_id" {
    name = "/terraform/iam/terraform-user-access-key-id"
    value     = "${aws_iam_access_key.terraform.id}"
    type      = "SecureString"
    overwrite = true
}

resource "aws_ssm_parameter" "terraform_user_secret_access_key" {
    name = "/terraform/iam/terraform-user-secret-access-key"
    value     = "${aws_iam_access_key.terraform.secret}"
    type      = "SecureString"
    overwrite = true
}

resource "aws_ssm_parameter" "service-lambda_role_arn" {
    name = "/terraform/iam/service-lambda-role-arn"
    type = "String"
    value = aws_iam_role.service.arn
}

resource "aws_ssm_parameter" "service_chat_lambda_role_arn" {
    name = "/terraform/iam/service-chat-lambda-role-arn"
    type = "String"
    value = aws_iam_role.service_chat.arn
}
