resource "aws_dynamodb_table" "booking" {
  billing_mode = "PAY_PER_REQUEST"
  name = "${var.booking_table_name}-${var.environment}"
  hash_key = "UId"
  stream_enabled = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "UId"
    type = "S"
  }

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "SpaceId"
    type = "S"
  }

  attribute {
    name = "BookingStatus"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  global_secondary_index {
    name = "BookingSpaceIdCreatedAtIndex"
    hash_key = "SpaceId"
    range_key = "CreatedAt"
    projection_type = "ALL"
  }

  global_secondary_index {
    name = "BookingSpaceIdBookingStatusIndex"
    hash_key = "SpaceId"
    range_key = "BookingStatus"
    projection_type = "ALL"
  }

  global_secondary_index {
    name = "BookingUserIdCreatedAtIndex"
    hash_key = "UserId"
    range_key = "CreatedAt"
    projection_type = "ALL"
  }

  global_secondary_index {
    name = "BookingUserIdBookingStatusIndex"
    hash_key = "UserId"
    range_key = "BookingStatus"
    projection_type = "ALL"
  }
}