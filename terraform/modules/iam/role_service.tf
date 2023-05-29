resource "aws_iam_role_policy_attachment" "enable_service_log" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_log.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_dynamodb_stream" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_dynamodb_stream.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_dynamodb" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_dynamodb.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_ssm" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_ssm.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_sqs" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_sqs.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_sns" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_sns.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_s3" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_s3.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_kinesis" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_kinesis.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_vpc_access" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_vpc_access.arn
}

resource "aws_iam_role_policy_attachment" "enable_service_elasticache" {
  role       = aws_iam_role.service.name
  policy_arn = aws_iam_policy.enable_elasticache.arn
}

resource "aws_iam_role" "service" {
  name                  = "service-lambda-role"
  force_detach_policies = true

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}