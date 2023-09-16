resource "aws_vpc_endpoint" "ecr_api" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.ecr.api"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "ecr-api"
  }
}

resource "aws_vpc_endpoint" "ecr_dkr" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.ecr.dkr"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "ecr-dkr"
  }
}

resource "aws_vpc_endpoint" "elasticache" {
  vpc_id = aws_vpc.this.id
  service_name = "com.amazonaws.us-east-1.elasticache"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true
  
  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "elasticache"
  }
}

resource "aws_vpc_endpoint" "rds" {
  vpc_id            = aws_vpc.this.id
  service_name      = "com.amazonaws.us-east-1.rds"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "rds"
  }
}

resource "aws_vpc_endpoint" "sqs" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.sqs"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "sqs"
  }
}

resource "aws_vpc_endpoint" "sns" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.sns"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "sns"
  }
}

resource "aws_vpc_endpoint" "kinesis" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.kinesis-streams"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "kinesis"
  }
}

resource "aws_vpc_endpoint" "logs" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.logs"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = true

  security_group_ids = [ aws_security_group.vpc_endpoint.id ]
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "logs"
  }
}

resource "aws_vpc_endpoint" "dynamodb" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.dynamodb"
  vpc_endpoint_type = "Gateway"

  tags = {
    Name = "dynamodb"
  }
}

resource "aws_vpc_endpoint" "s3" {
  vpc_id       = aws_vpc.this.id
  service_name = "com.amazonaws.${var.region}.s3"
  vpc_endpoint_type = "Gateway"

  tags = {
    Name = "s3"
  }
}
