resource "aws_flow_log" "this" {
  vpc_id          = aws_vpc.this.id
  iam_role_arn    = var.flow_logs_role
  traffic_type    = "ALL"
  log_destination = "arn:aws:logs:us-east-1:${var.account_id}:log-group:flow-logs"
}
