resource "aws_ssm_parameter" "vpc" {
  name = "/terraform/ec2/vpc-id"
  description = "VPC Id provided by terraform"
  type = "String"
  value = aws_vpc.default.id
}