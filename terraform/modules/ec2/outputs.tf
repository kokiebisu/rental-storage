output "serverless_security_group_id" {
    value = aws_security_group.serverless.id
}

output "db_subnet_group_name" {
    value = aws_db_subnet_group.serverless.name
}