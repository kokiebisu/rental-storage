resource "aws_dynamodb_table" "connections" {
  billing_mode = "PAY_PER_REQUEST"
  name         = "${var.environment}-${var.chat_db_info.table_name}-connections"
  hash_key     = "ConnectionId"

  attribute {
    name = "ConnectionId"
    type = "S"
  }

  attribute {
    name = "SenderId"
    type = "S"
  }

  attribute {
    name = "RecipientId"
    type = "S"
  }

  global_secondary_index {
    name            = "ChatConnectionsSenderIdRecipientIdIndex"
    hash_key        = "SenderId"
    range_key       = "RecipientId"
    projection_type = "ALL"
  }
}

resource "aws_dynamodb_table" "messages" {
  billing_mode     = "PAY_PER_REQUEST"
  name             = "${var.environment}-${var.chat_db_info.table_name}-messages"
  hash_key         = "Id"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "Id"
    type = "S"
  }

  attribute {
    name = "SenderId"
    type = "S"
  }

  attribute {
    name = "RecipientId"
    type = "S"
  }

  global_secondary_index {
    name            = "ChatMessagesSenderIdRecipientIdIndex"
    hash_key        = "SenderId"
    range_key       = "RecipientId"
    projection_type = "ALL"
  }
}
