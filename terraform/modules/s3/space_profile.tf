resource "aws_s3_bucket" "space_profile" {
  bucket = "${var.environment}-${var.account_id}-space-profile"
}

# resource "aws_s3_bucket_public_access_block" "space_profile" {
#   bucket = aws_s3_bucket.space_profile.id

#   block_public_acls       = true
#   block_public_policy     = false
#   ignore_public_acls      = true
#   restrict_public_buckets = true
# }

resource "aws_s3_bucket_cors_configuration" "space_profile" {
  bucket = aws_s3_bucket.space_profile.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT", "POST", "HEAD"]
    allowed_origins = ["*"]
  }
}
