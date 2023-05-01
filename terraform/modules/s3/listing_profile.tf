# resource "aws_s3_bucket" "space_profile" {
#   bucket = "${var.environment}-${var.account_id}-space-profile"
# }

# resource "aws_s3_bucket_cors_configuration" "space_profile" {
#   bucket = aws_s3_bucket.space_profile.id

#   cors_rule {
#     allowed_methods = ["GET", "PUT", "POST", "HEAD"]
#     allowed_origins = ["*"]
#     allowed_headers = ["*"]
#   }
# }

# resource "aws_s3_bucket_acl" "space_profile" {
#   bucket = aws_s3_bucket.space_profile.id
#   acl    = "public-read-write"
# }

# resource "aws_s3_bucket_policy" "space_profile" {
#   bucket = aws_s3_bucket.space_profile.id
#   policy = data.aws_iam_policy_document.space_profile.json
# }

# data "aws_iam_policy_document" "space_profile" {
#   statement {
#     sid    = "PublicReadForGetBucketObjects"
#     effect = "Allow"

#     principals {
#       type        = "*"
#       identifiers = ["*"]
#     }

#     actions = ["s3:GetObject", "s3:PutObject"]

#     resources = ["${aws_s3_bucket.space_profile.arn}/*"]
#   }
# }