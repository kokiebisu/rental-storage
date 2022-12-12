module "dynamodb" {
  source = "./modules/dynamodb"

  booking_table_name = var.booking_table_name
  environment        = var.environment
}

module "ec2" {
  source = "./modules/ec2"

  namespace = var.namespace
}

module "iam" {
  source = "./modules/iam"

  region     = var.region
  account_id = module.identity.account_id
}

module "identity" {
  source = "./modules/identity"
}

module "kinesis" {
  source = "./modules/kinesis"

  environment = var.environment
}

module "rds" {
  source = "./modules/rds"

  namespace                    = var.namespace
  environment                  = var.environment
  serverless_security_group_id = module.ec2.serverless_security_group_id
  db_subnet_group_name         = module.ec2.db_subnet_group_name

  listing_db_username = var.listing_db_username
  listing_db_password = var.listing_db_password
  listing_db_name     = var.listing_db_name

  user_db_username = var.user_db_username
  user_db_password = var.user_db_password
  user_db_name     = var.user_db_name
}

module "s3" {
  source = "./modules/s3"

  namespace   = var.namespace
  environment = var.environment
  account_id  = module.identity.account_id
}

module "sqs" {
  source = "./modules/sqs"

  namespace                = var.namespace
  environment              = var.environment
  authentication_topic_arn = module.sns.authentication_topic_arn
  booking_topic_arn        = module.sns.booking_topic_arn
  listing_topic_arn        = module.sns.listing_topic_arn
  payment_topic_arn        = module.sns.payment_topic_arn
  user_topic_arn           = module.sns.user_topic_arn
}

module "sns" {
  source = "./modules/sns"

  namespace         = var.namespace
  environment       = var.environment
  booking_queue_arn = module.sqs.booking_queue_arn
  listing_queue_arn = module.sqs.listing_queue_arn
  payment_queue_arn = module.sqs.payment_queue_arn
  user_queue_arn    = module.sqs.user_queue_arn
}
