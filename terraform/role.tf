resource "aws_iam_policy" "enable_log" {
    name = "EnableLog"
    description = "Enable logs"

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

resource "aws_iam_policy" "enable_dynamodb_stream" {
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
                Resource = "arn:aws:dynamodb:${var.region}:${module.identity.account_id}:table/*/stream/*"
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
                Action = ["dynamodb:PutItem", "dynamodb:UpdateItem", "dynamodb:DeleteItem", "dynamodb:GetItem", "dynamodb:Scan", "dynamodb:Query"]
                Effect = "Allow"
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
                Action = ["ssm:GetParameter", "ssm:GetParameters", "ssm:GetParametersByPath", "ssm:DescribeParameters", "ssm:PutParameter"]
                Effect = "Allow"
                Resource = "arn:aws:ssm:${var.region}:${module.identity.account_id}:*"
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
                Action = ["sqs:ReceiveMessage", "sqs:DeleteMessage", "sqs:GetQueueAttributes"]
                Effect = "Allow"
                Resource = "arn:aws:sqs:${var.region}:${module.identity.account_id}:*"
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
                Action = ["sns:CreateTopic", "sns:Publish"],
                Effect = "Allow"
                Resource = "arn:aws:sns:${var.region}:${module.identity.account_id}:*"
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
                Action = ["s3:*"],
                Effect = "Allow"
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
                Action = ["kinesis:PutRecord", "kinesis:PutRecords"]
                Effect = "Allow"
                Resource = "arn:aws:kinesis:${var.region}:${module.identity.account_id}:stream/*"
            },
            {
                Action = ["kinesis:Get*", "kinesis:List*", "kinesis:Describe*"]
                Effect = "Allow"
                Resource = "arn:aws:kinesis:${var.region}:${module.identity.account_id}:stream/*"
            }
        ]
    })
}

resource "aws_iam_role_policy_attachment" "enable_log" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_log.arn
}

resource "aws_iam_role_policy_attachment" "enable_dynamodb_stream" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_dynamodb_stream.arn
}

resource "aws_iam_role_policy_attachment" "enable_dynamodb" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_dynamodb.arn
}

resource "aws_iam_role_policy_attachment" "enable_ssm" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_ssm.arn
}

resource "aws_iam_role_policy_attachment" "enable_sqs" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_sqs.arn
}

resource "aws_iam_role_policy_attachment" "enable_sns" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_sns.arn
}

resource "aws_iam_role_policy_attachment" "enable_s3" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_s3.arn
}

resource "aws_iam_role_policy_attachment" "enable_kinesis" {
    role = aws_iam_role.lambda.name
    policy_arn = aws_iam_policy.enable_kinesis.arn
}

resource "aws_iam_role" "lambda" {
    name = "service-lambda-role"
    force_detach_policies = true
    
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
}

resource "aws_ssm_parameter" "lambda_role_arn" {
    name = "/terraform/iam/lambda-role-arn"
    type = "String"
    value = aws_iam_role.lambda.arn
}