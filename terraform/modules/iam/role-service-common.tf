resource "aws_iam_role_policy_attachment" "flowlogs_policy_attachment" {
  policy_arn = aws_iam_policy.enable_flow_logs.arn
  role       = aws_iam_role.flowlogs_role.name
}

resource "aws_iam_role" "flowlogs_role" {
  name = "flow-logs-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "vpc-flow-logs.amazonaws.com"
        }
      }
    ]
  })
}
