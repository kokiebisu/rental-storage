resource "aws_dynamodb_table" "connections" {
  billing_mode = "PAY_PER_REQUEST"
  name = "${var.environment}-${var.chat_table_name}-connections"
  hash_key = "ConnectionId"

  attribute {
    name = "ConnectionId"
    type = "S"
  }

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "DestinationUserId"
    type = "S"
  }

  global_secondary_index {
    name = "ChatUserIdDestinationIdIndex"
    hash_key = "UserId"
    range_key = "DestinationUserId"
    projection_type = "ALL"
  }
}

resource "aws_dynamodb_table" "messages" {
  billing_mode = "PAY_PER_REQUEST"
  name = "${var.environment}-${var.chat_table_name}-messages"
  hash_key = "UId"
  stream_enabled = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "UId"
    type = "S"
  }

  attribute {
    name = "GuestId"
    type = "S"
  }

  attribute {
    name = "HostId"
    type = "S"
  }

  global_secondary_index {
    name = "ChatGuestIdHostIdIndex"
    hash_key = "GuestId"
    range_key = "HostId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name = "ChatHostIdGuestIdIndex"
    hash_key = "HostId"
    range_key = "GuestId"
    projection_type = "ALL"
  }
}