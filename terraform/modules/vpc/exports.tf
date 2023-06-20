resource "aws_ssm_parameter" "vpc" {
  name        = "/terraform/vpc/vpc-id"
  description = "VPC Id provided by terraform"
  type        = "String"
  value       = aws_vpc.this.id
}

resource "aws_ssm_parameter" "subnet_a_id" {
  name        = "/terraform/vpc/subnet-a-id"
  description = "Subnet A provided by terraform"
  type        = "String"
  value       = aws_subnet.a.id
}

resource "aws_ssm_parameter" "subnet_b_id" {
  name        = "/terraform/vpc/subnet-b-id"
  description = "Subnet B provided by terraform"
  type        = "String"
  value       = aws_subnet.b.id
}

resource "aws_ssm_parameter" "subnet_c_id" {
  name        = "/terraform/vpc/subnet-c-id"
  description = "Subnet C provided by terraform"
  type        = "String"
  value       = aws_subnet.c.id
}

resource "aws_ssm_parameter" "subnet_d_id" {
  name        = "/terraform/vpc/subnet-d-id"
  description = "Subnet D provided by terraform"
  type        = "String"
  value       = aws_subnet.d.id
}

resource "aws_ssm_parameter" "lambda_security_group" {
  name        = "/terraform/vpc/lambda-security-group"
  description = "Lambda Security Group provided by terraform"
  type        = "String"
  value       = aws_security_group.lambda.id
}

resource "aws_ssm_parameter" "ec2_security_group" {
  name        = "/terraform/vpc/ec2-security-group"
  description = "EC2 Security Group provided by terraform"
  type        = "String"
  value       = aws_security_group.ec2.id
}

resource "aws_ssm_parameter" "alb_security_group" {
  name        = "/terraform/vpc/alb-security-group"
  description = "Application Load Balancer Security Group provided by terraform"
  type        = "String"
  value       = aws_security_group.alb.id
}

resource "aws_ssm_parameter" "rds_security_group" {
  name        = "/terraform/vpc/rds-security-group"
  description = "RDS Security Group provided by terraform"
  type        = "String"
  value       = aws_security_group.rds_postgres.id
}

resource "aws_ssm_parameter" "rds_subnet_group" {
  name        = "/terraform/vpc/rds-subnet-group"
  description = "RDS Subnet Group provided by terraform"
  type        = "String"
  value       = aws_db_subnet_group.this.id
}

resource "aws_ssm_parameter" "elasticache_subnet_group" {
  name        = "/terraform/vpc/elasticache-subnet-group"
  description = "Elasticache Subnet Group provided by terraform"
  type        = "String"
  value       = aws_elasticache_subnet_group.this.id
}
