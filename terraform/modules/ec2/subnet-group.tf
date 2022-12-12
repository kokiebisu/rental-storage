resource "aws_db_subnet_group" "serverless" {
  description = "RDS Subnet Group"
  subnet_ids = [ aws_subnet.a.id, aws_subnet.b.id, aws_subnet.c.id ]

  tags = {
    Name = "DBSubnetGroupServerless"
  }
}