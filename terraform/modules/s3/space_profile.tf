resource "aws_s3_bucket" "space_profile" {
  bucket = "${var.environment}-${var.account_id}-space-profile"
}

resource "aws_iam_policy" "s3" {
  name        = "example-policy"
  description = "Allows for uploading files to Example S3 bucket using getPresignedUrl"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "ExampleSid",
        Effect    = "Allow",
        Principal = {
          AWS = "${var.aws_iam_user_s3_arn}"
        },
        Action    = [
          "s3:PutObject",
          "s3:GetObject"
        ],
        Resource  = [
          "arn:aws:s3:::example-bucket/*",
        ]
      }
    ]
  })
}

resource "aws_iam_policy_attachment" "s3" {
  name       = "example-policy-attachment"
  policy_arn = aws_iam_policy.s3.arn
  users      = [var.aws_iam_user_s3_name]
}
