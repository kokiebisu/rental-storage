output "vpc_id" {
    value = aws_vpc.this.id
}

output "lambda_security_group_id" {
    value = aws_security_group.lambda.id
}

output "ec2_security_group_id" {
    value = aws_security_group.ec2.id
}

output "rds_postgres_security_group_id" {
    value = aws_security_group.rds_postgres.id
}

output "rds_subnet_group_name" {
    value = aws_db_subnet_group.this.name
}

output "elasticache_security_group_id" {
    value = aws_security_group.elasticache.id
}

output "elasticache_subnet_group_name" {
    value = aws_elasticache_subnet_group.this.name
}

output "elasticache_preferred_availability_zone" {
    value = aws_subnet.b.id
}

output "primary_public_subnet_id" {
    value = aws_subnet.a.id
}
