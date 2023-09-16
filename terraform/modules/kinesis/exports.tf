resource "aws_ssm_parameter" "space_db_host" {
  name  = "/terraform/kinesis/event-stream-arn"
  type  = "String"
  value = aws_kinesis_stream.event.arn
} 