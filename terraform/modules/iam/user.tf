resource "aws_iam_user" "service_image_presigned" {
  name = "service-image-presigned"
}

resource "aws_iam_access_key" "service_image_presigned" {
  user = aws_iam_user.service_image_presigned.name
}

resource "aws_iam_policy_attachment" "service_image_presigned-attach" {
  name = "service-image-presigned-policy-attachment"
  users      = [aws_iam_user.service_image_presigned.name]
  policy_arn = aws_iam_policy.enable_s3.arn
}
