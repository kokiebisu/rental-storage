resource "aws_s3_bucket_policy" "listing_profile" {
  bucket = aws_s3_bucket.listing_profile.id
  policy = data.aws_iam_policy_document.listing_profile.json
}

data "aws_iam_policy_document" "listing_profile" {
    statement {
        sid = "PublicReadForGetBucketObjects"
        effect = "Allow"

        principals {
          type = "*"
          identifiers = ["*"]
        }

        actions = [ "s3:GetObject", "s3:PutObject" ]

        resources = [ "${aws_s3_bucket.listing_profile.arn}/*" ]
    }
}