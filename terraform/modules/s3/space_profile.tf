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
