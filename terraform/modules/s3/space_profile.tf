resource "aws_s3_bucket" "space_profile" {
  bucket = "${var.environment}-${var.account_id}-space-profile"

  force_destroy = true
}

resource "aws_s3_bucket_cors_configuration" "space_profile" {
  bucket = aws_s3_bucket.space_profile.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT", "POST", "HEAD"]
    allowed_origins = ["*"]
  }
}

resource "aws_s3_bucket_public_access_block" "example" {
  bucket = aws_s3_bucket.space_profile.id

  block_public_acls       = true
  block_public_policy     = false
  ignore_public_acls      = true
  restrict_public_buckets = false
}

resource "aws_s3_bucket_policy" "space_profile" {
  bucket = aws_s3_bucket.space_profile.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = "*"
        Action = "s3:GetObject"
        Resource = "${aws_s3_bucket.space_profile.arn}/*"
      }
    ]
  })
}
