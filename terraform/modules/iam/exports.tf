resource "aws_ssm_parameter" "service_image_presigned_user_access_key_id" {
    name = "/terraform/iam/service-image-presigned-user-access-key-id"
    value     = "${aws_iam_access_key.service_image_presigned.id}"
    type      = "SecureString"
    overwrite = true
}

resource "aws_ssm_parameter" "service_image_presigned_user_secret_access_key" {
    name = "/terraform/iam/service-image-presigned-user-secret-access-key"
    value     = "${aws_iam_access_key.service_image_presigned.secret}"
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
