resource "aws_s3_bucket" "listing_profile" {
  bucket = "${var.environment}-${var.account_id}-listing-profile"
  
}

resource "aws_s3_bucket_cors_configuration" "listing_profile" {
  bucket = aws_s3_bucket.listing_profile.id

  cors_rule {
    allowed_methods = ["GET", "PUT", "POST", "HEAD"]
    allowed_origins = [ "*" ]
    allowed_headers = [ "*" ]
  }
}

resource "aws_s3_bucket_acl" "listing_profile" {
  bucket = aws_s3_bucket.listing_profile.id
  acl = "public-read-write"
}