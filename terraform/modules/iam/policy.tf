resource "aws_iam_policy" "enable_log" {
  name        = "EnableLog"
  description = "Enable logs"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["logs:CreateLogGroup", "logs:CreateLogStream", "logs:PutLogEvents"]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_dynamodb_stream" {
  name = "EnableDynamoDBStream"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["dynamodb:ListStreams"]
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Action   = ["dynamodb:GetRecords", "dynamodb:DescribeStream", "dynamodb:GetShardIterator"]
        Effect   = "Allow"
        Resource = "arn:aws:dynamodb:${var.region}:${var.account_id}:table/*/stream/*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_dynamodb" {
  name = "EnableDynamoDB"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["dynamodb:PutItem", "dynamodb:UpdateItem", "dynamodb:DeleteItem", "dynamodb:GetItem", "dynamodb:Scan", "dynamodb:Query"]
        Effect   = "Allow"
        Resource = "arn:aws:dynamodb:${var.region}:*:table/*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_ssm" {
  name = "EnableSSM"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["ssm:GetParameter", "ssm:GetParameters", "ssm:GetParametersByPath", "ssm:DescribeParameters", "ssm:PutParameter"]
        Effect   = "Allow"
        Resource = "arn:aws:ssm:${var.region}:${var.account_id}:*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_sqs" {
  name = "EnableSQS"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["sqs:ReceiveMessage", "sqs:DeleteMessage", "sqs:GetQueueAttributes"]
        Effect   = "Allow"
        Resource = "arn:aws:sqs:${var.region}:${var.account_id}:*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_sns" {
  name = "EnableSNS"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["sns:CreateTopic", "sns:Publish"],
        Effect   = "Allow"
        Resource = "arn:aws:sns:${var.region}:${var.account_id}:*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_s3" {
  name = "EnableS3"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["s3:*"],
        Effect   = "Allow"
        Resource = "arn:aws:s3:::*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_kinesis" {
  name = "EnableKinesis"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["kinesis:PutRecord", "kinesis:PutRecords"]
        Effect   = "Allow"
        Resource = "arn:aws:kinesis:${var.region}:${var.account_id}:stream/*"
      },
      {
        Action   = ["kinesis:Get*", "kinesis:List*", "kinesis:Describe*"]
        Effect   = "Allow"
        Resource = "arn:aws:kinesis:${var.region}:${var.account_id}:stream/*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_vpc_access" {
  name = "EnableVPCAccess"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface",
          "ec2:AssignPrivateIpAddresses",
          "ec2:UnassignPrivateIpAddresses"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_elasticache" {
  name = "EnableElastiCache"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["elasticache:Describe*", "elasticache:List*", "elasticache:Create*", "elasticache:Modify*", "elasticache:Delete*", "elasticache:AuthorizeCacheSecurityGroupIngress"]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_policy" "enable_api_gateway_invoke" {
  name = "EnableAPIGatewayInvoke"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["execute-api:Invoke", "execute-api:ManageConnections"]
        Effect   = "Allow"
        Resource = "arn:aws:execute-api:*:*:*"
      }
    ]
  })
}
