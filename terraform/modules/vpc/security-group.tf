resource "aws_security_group" "any" {
  description = "Security group for Serverless functions"
  vpc_id = aws_vpc.this.id
  
  tags = {
    Name = "SecurityGroupServerless"
  }
}

resource "aws_security_group" "elasticache" {
  description = "Security group for Elasticache"
  vpc_id = aws_vpc.this.id

  tags = {
    Name = "SecurityGroupElasticache"
  }
}

resource "aws_security_group" "ec2" {
  description = "Security group for EC2"
  vpc_id = aws_vpc.this.id
}

resource "aws_security_group_rule" "allow_any" {
  type        = "ingress"
  from_port   = 0
  to_port     = 65535
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.any.id
}

resource "aws_security_group_rule" "allow_elasticache_redis" {
  type        = "ingress"
  from_port   = 6379
  to_port     = 6379
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.elasticache.id
}

resource "aws_security_group_rule" "allow_ssh" {
  type        = "ingress"
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.ec2.id
}

resource "aws_security_group_rule" "allow_ec2_elasticache_redis" {
  type        = "egress"
  from_port   = 6379
  to_port     = 6379
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.ec2.id
}

resource "aws_security_group_rule" "allow_lambda_elasticache_redis" {
  type        = "egress"
  from_port   = 6379
  to_port     = 6379
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.any.id
}

resource "aws_security_group_rule" "allow_lambda_apt_traffic" {
  type        = "egress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"] # Replace with your desired CIDR block

  security_group_id = aws_security_group.any.id # Replace with the ID of your existing security group
}

resource "aws_security_group_rule" "allow_lambda_apt_traffic_https" {
  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
  
  security_group_id = aws_security_group.any.id # Replace with the ID of your existing security group
}

resource "aws_security_group_rule" "allow_apt_traffic" {
  type        = "egress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"] # Replace with your desired CIDR block

  security_group_id = aws_security_group.ec2.id # Replace with the ID of your existing security group
}

resource "aws_security_group_rule" "allow_apt_traffic_https" {
  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
  
  security_group_id = aws_security_group.ec2.id # Replace with the ID of your existing security group
}
