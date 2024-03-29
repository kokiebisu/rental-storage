module "apigateway" {
  source = "./modules/apigateway"

  namespace   = var.namespace
  environment = var.environment
}

module "cloudwatch_logs" {
  source = "./modules/cloudwatch_logs"
}

module "dynamodb" {
  source = "./modules/dynamodb"

  namespace   = var.namespace
  environment = var.environment

  booking_db_info = var.booking_db_info
  chat_db_info    = var.chat_db_info
}

module "ecr" {
  source = "./modules/ecr"

  namespace   = var.namespace
  environment = var.environment
}

module "ec2" {
  source = "./modules/ec2"

  namespace   = var.namespace
  environment = var.environment

  vpc_id                       = module.vpc.vpc_id
  public_ec2_security_group_id = module.vpc.public_ec2_security_group_id
  primary_public_subnet_id     = module.vpc.primary_public_subnet_id
  secondary_public_subnet_id   = module.vpc.secondary_public_subnet_id
}

module "elasticache" {
  source = "./modules/elasticache"

  namespace   = var.namespace
  environment = var.environment

  elasticache_security_group_id           = module.vpc.elasticache_security_group_id
  elasticache_subnet_group_name           = module.vpc.elasticache_subnet_group_name
  elasticache_preferred_availability_zone = module.vpc.elasticache_preferred_availability_zone
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

  namespace   = var.namespace
  environment = var.environment
}

module "rds" {
  source = "./modules/rds"

  namespace   = var.namespace
  environment = var.environment

  db_security_group_id = module.vpc.rds_postgres_security_group_id
  db_subnet_group_name = module.vpc.rds_subnet_group_name

  user_db_info  = var.user_db_info
  space_db_info = var.space_db_info

  #   space_db_username    = var.space_db_username
  #   space_db_password    = var.space_db_password
  #   space_db_name        = var.space_db_name
  #   user_db_username     = var.user_db_username
  #   user_db_password     = var.user_db_password
  #   user_db_name         = var.user_db_name
}

module "s3" {
  source = "./modules/s3"

  namespace   = var.namespace
  environment = var.environment

  account_id = module.identity.account_id
}

module "sqs" {
  source = "./modules/sqs"

  namespace   = var.namespace
  environment = var.environment

  authentication_topic_arn = module.sns.authentication_topic_arn
  booking_topic_arn        = module.sns.booking_topic_arn
  space_topic_arn          = module.sns.space_topic_arn
  payment_topic_arn        = module.sns.payment_topic_arn
  user_topic_arn           = module.sns.user_topic_arn
}

module "sns" {
  source = "./modules/sns"

  namespace   = var.namespace
  environment = var.environment

  booking_queue_arn = module.sqs.booking_queue_arn
  space_queue_arn   = module.sqs.space_queue_arn
  payment_queue_arn = module.sqs.payment_queue_arn
  user_queue_arn    = module.sqs.user_queue_arn
}

module "vpc" {
  source = "./modules/vpc"

  region         = var.region
  account_id     = module.identity.account_id
  flow_logs_role = module.iam.flow_logs_role
  public_ec2_id  = module.ec2.public_ec2_id
}
