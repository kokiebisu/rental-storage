# need to create iam 'terraform' user first
resource "aws_s3_bucket_policy" "space_profile" {
  bucket = aws_s3_bucket.space_profile.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
          Sid = "AllowAllS3Actions",
          Effect = "Allow",
          Action = [
              "s3:GetObject",
              "s3:PutObject",
              "s3:PutObjectAcl"
          ],
          Principal = "*",
          Resource = "${aws_s3_bucket.space_profile.arn}/*"
      }
    ]
  })
}
