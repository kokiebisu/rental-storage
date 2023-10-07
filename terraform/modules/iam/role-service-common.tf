resource "aws_iam_role_policy_attachment" "flowlogs_policy_attachment" {
  policy_arn = aws_iam_policy.enable_flow_logs.arn
  role       = aws_iam_role.flowlogs_role.name
}

resource "aws_iam_role" "flowlogs_role" {
  name               = "flow-logs-role"
  assume_role_policy = file("${local.path}/role-flow-logs.tpl")
}
