resource "aws_iam_role" "lambda" {
    name = "service-lambda-role"
    assume_role_policy = jsonencode({
        Version = "2012-10-17"
        Statement = [
            {
                Action = "sts:AssumeRole"
                Effect = "Allow"
                Sid = ""
                Principal = {
                    Service = "lambda.amazonaws.com"
                }
            }
        ]
    })

    inline_policy {
      name = "EnableLog"

      policy = jsonencode({
        Version = "2012-10-17"
        Statement = [
            {
                Action = ["logs:CreateLogGroup", "logs:CreateLogStream", "logs:PutLogEvents"]
                Effect = "Allow"
                Resource = "arn:aws:logs:*:*:*"
            }
        ]
      })
    }

    inline_policy {
      name = "EnableDynamoDBStream"

      policy = jsonencode({
        Version = "2012-10-17"
        Statement = [
            {
                Action = ["dynamodb:ListStreams"]
                Effect = "Allow"
                Resource = "*"
            },
            {
                Action = ["dynamodb:GetRecords", "dynamodb:DescribeStream", "dynamodb:GetShardIterator"]
                Effect = "Allow"
                Resource = "arn:aws:dynamodb:${var.region}:${var.account_id}:table/*/stream/*"
            }
        ]
      })
    }

    inline_policy {
        name = "EnableDynamoDB"

        policy = jsonencode({
            Version = "2012-10-17"
            Statement = [
                {
                    Action = ["dynamodb:PutItem", "dynamodb:UpdateItem", "dynamodb:DeleteItem", "dynamodb:GetItem", "dynamodb:Scan", "dynamodb:Query"]
                    Effect = "Allow"
                    Resource = "arn:aws:dynamodb:${var.region}:*:table/*"
                }
            ]
        })
    }

    inline_policy {
      name = "EnableSSM"

      policy = jsonencode({
        Version = "2012-10-17"
        Statement = [
            {
                Action = ["ssm:GetParameter", "ssm:GetParameters", "ssm:GetParametersByPath", "ssm:DescribeParameters", "ssm:PutParameter"]
                Effect = "Allow"
                Resource = "arn:aws:ssm:${var.region}:${var.account_id}:*"
            }
        ]
      })
    }

    inline_policy {
        name = "EnableSQS"

        policy = jsonencode({
            Version = "2012-10-17"
            Statement = [
                {
                    Action = ["sqs:ReceiveMessage", "sqs:DeleteMessage", "sqs:GetQueueAttributes"]
                    Effect = "Allow"
                    Resource = "arn:aws:sqs:${var.region}:${var.account_id}:*"
                }
            ]
        })
    }

    inline_policy {
        name = "EnableSNS"

        policy = jsonencode({
            Version = "2012-10-17"
            Statement = [
                {
                    Action = ["sns:CreateTopic", "sns:Publish"],
                    Effect = "Allow"
                    Resource = "arn:aws:sns:${var.region}:${var.account_id}:*"
                }
            ]
        })
    }

    inline_policy {
        name = "EnableS3"

        policy = jsonencode({
            Version = "2012-10-17"
            Statement = [
                {
                    Action = ["s3:*"],
                    Effect = "Allow"
                    Resource = "arn:aws:s3:::*"
                }
            ]
        })
    }

    inline_policy {
        name = "EnableKinesis"

        policy = jsonencode({
            Version = "2012-10-17"
            Statement = [
                {
                    Action = ["kinesis:PutRecord", "kinesis:PutRecords"]
                    Effect = "Allow"
                    Resource = "arn:aws:kinesis:${var.region}:${var.account_id}:stream/*"
                },
                {
                    Action = ["kinesis:Get*", "kinesis:List*", "kinesis:Describe*"]
                    Effect = "Allow"
                    Resource = "arn:aws:kinesis:${var.region}:${var.account_id}:stream/*"
                }
            ]
        })
    }
}
