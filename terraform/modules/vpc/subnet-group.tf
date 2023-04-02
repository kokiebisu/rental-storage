resource "aws_db_subnet_group" "serverless" {
  description = "RDS Subnet Group"
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "db-subnet-group"
  }
}

resource "aws_elasticache_subnet_group" "serverless" {
  description       = "Elasticache Subnet Group"
  name = "elasticache-subnet-group"
  subnet_ids = [ aws_subnet.b.id, aws_subnet.d.id ]

  tags = {
    Name = "elasticache-subnet-group"
  }
}
