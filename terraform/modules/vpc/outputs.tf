output "vpc_id" {
    value = aws_vpc.this.id
}

output "serverless_security_group_id" {
    value = aws_security_group.any.id
}

output "rds_postgres_security_group_id" {
    value = aws_security_group.rds_postgres.id
}

output "rds_subnet_group_name" {
    value = aws_db_subnet_group.serverless.name
}

output "elasticache_security_group_id" {
    value = aws_security_group.elasticache.id
}

output "elasticache_subnet_group_name" {
    value = aws_elasticache_subnet_group.serverless.name
}

output "elasticache_preferred_availability_zone" {
    value = aws_subnet.a.id
}

output "primary_subnet_id" {
  value = aws_subnet.a.id
}
