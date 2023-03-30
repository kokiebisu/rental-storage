resource "aws_ssm_parameter" "vpc" {
  name = "/terraform/ec2/vpc-id"
  description = "VPC Id provided by terraform"
  type = "String"
  value = aws_vpc.this.id
}

resource "aws_ssm_parameter" "serverless-security-group" {
  name = "/terraform/ec2/security-group-id"
  description = "Security Group Id provided by terraform"
  type = "String"
  value = aws_security_group.any.id
}

resource "aws_ssm_parameter" "elasticache-security-group" {
  name = "/terraform/ec2/elasticache-security-group-id"
  description = "Elasticache Security Group Id provided by terraform"
  type = "String"
  value = aws_security_group.elasticache.id
}

resource "aws_ssm_parameter" "subnet-a" {
  name = "/terraform/ec2/subnet-a-id"
  description = "Subnet A Id provided by terraform"
  type = "String"
  value = aws_subnet.a.id
}

resource "aws_ssm_parameter" "subnet-b" {
  name = "/terraform/ec2/subnet-b-id"
  description = "Subnet B Id provided by terraform"
  type = "String"
  value = aws_subnet.b.id
}

resource "aws_ssm_parameter" "subnet-c" {
  name = "/terraform/ec2/subnet-c-id"
  description = "Subnet C Id provided by terraform"
  type = "String"
  value = aws_subnet.c.id
}
