locals {
  resource_input = {
    region = "${var.region}",
    account_id = "${var.account_id}"
  }
}

resource "aws_iam_policy" "enable_log" {
  name        = "EnableLog"
  policy = file("${local.path}/policy-enable-log.tpl")
}

resource "aws_iam_policy" "enable_dynamodb_stream" {
  name = "EnableDynamoDBStream"
  policy = templatefile("${local.path}/policy-enable-dynamodb-stream.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_dynamodb" {
  name = "EnableDynamoDB"
  policy = templatefile("${local.path}/policy-enable-dynamodb.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_ssm" {
  name = "EnableSSM"
  policy = templatefile("${local.path}/policy-enable-ssm.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_sqs" {
  name = "EnableSQS"
  policy = templatefile("${local.path}/policy-enable-sqs.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_sns" {
  name = "EnableSNS"
  policy = templatefile("${local.path}/policy-enable-sns.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_s3" {
  name = "EnableS3"
  policy = file("${local.path}/policy-enable-s3.tpl")
}

resource "aws_iam_policy" "enable_kinesis" {
  name = "EnableKinesis"
  policy = templatefile("${local.path}/policy-enable-kinesis.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_vpc_access" {
  name = "EnableVPCAccess"
  policy = file("${local.path}/policy-enable-vpc-access.tpl")
}

resource "aws_iam_policy" "enable_elasticache" {
  name = "EnableElastiCache"
  policy = templatefile("${local.path}/policy-enable-elasticache.tpl", local.resource_input)
}

resource "aws_iam_policy" "enable_api_gateway_invoke" {
  name = "EnableAPIGatewayInvoke"
  policy = file("${local.path}/policy-enable-api-gateway-invoke.tpl")
}

resource "aws_iam_policy" "enable_flow_logs" {
  name        = "EnableFlowLogs"
  policy = file("${local.path}/policy-enable-flow-logs.tpl")
}
