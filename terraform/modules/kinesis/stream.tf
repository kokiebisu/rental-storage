resource "aws_kinesis_stream" "event" {
  name = "${var.environment}-EventStream"
  retention_period = 24
  shard_count = 1
}