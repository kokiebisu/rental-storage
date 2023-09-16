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

resource "aws_iam_policy" "flowlogs_policy" {
  name        = "flow-logs-policy"
  description = "Policy for VPC Flow Logs"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action   = "logs:CreateLogGroup",
        Effect   = "Allow",
        Resource = "arn:aws:logs:*:*:log-group:flow-logs"
      },
      {
        Action   = "logs:CreateLogStream",
        Effect   = "Allow",
        Resource = "arn:aws:logs:*:*:log-group:flow-logs:*"
      },
      {
        Action   = "logs:PutLogEvents",
        Effect   = "Allow",
        Resource = "arn:aws:logs:*:*:log-group:flow-logs:log-stream:*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "flowlogs_policy_attachment" {
  policy_arn = aws_iam_policy.flowlogs_policy.arn
  role       = aws_iam_role.flowlogs_role.name
}
