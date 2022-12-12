resource "aws_dynamodb_table" "booking" {
  billing_mode = "PAY_PER_REQUEST"
  name = "${var.booking_table_name}-${var.environment}"
  hash_key = "Id"
  stream_enabled = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "Id"
    type = "S"
  }

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "ListingId"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  global_secondary_index {
    name = "BookingUserIdIndex"
    hash_key = "UserId"
    range_key = "ListingId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name = "BookingCreatedAtIndex"
    hash_key = "UserId"
    range_key = "CreatedAt"
    projection_type = "ALL"
  }
}