resource "aws_ssm_parameter" "booking_table_name" {
  name  = "/terraform/dynamodb/booking-table-name"
  type  = "String"
  value = aws_dynamodb_table.booking.name
}

resource "aws_ssm_parameter" "booking_table_stream_arn" {
  name  = "/terraform/dynamodb/booking-table-stream-arn"
  type  = "String"
  value = aws_dynamodb_table.booking.stream_arn
}