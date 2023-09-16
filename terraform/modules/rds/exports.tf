resource "aws_ssm_parameter" "space_db_host" {
  name  = "/terraform/rds/space-db-host"
  type  = "String"
  value = aws_db_instance.space.address
}

resource "aws_ssm_parameter" "space_db_port" {
  name  = "/terraform/rds/space-db-port"
  type  = "String"
  value = aws_db_instance.space.port
}

resource "aws_ssm_parameter" "space_db_name" {
  name  = "/terraform/rds/space-db-name"
  type  = "String"
  value = aws_db_instance.space.db_name
}

resource "aws_ssm_parameter" "space_db_username" {
  name  = "/terraform/rds/space-db-username"
  type  = "String"
  value = aws_db_instance.space.username
}

resource "aws_ssm_parameter" "space_db_password" {
  name  = "/terraform/rds/space-db-password"
  type  = "String"
  value = aws_db_instance.space.password
}

resource "aws_ssm_parameter" "user_db_host" {
  name  = "/terraform/rds/user-db-host"
  type  = "String"
  value = aws_db_instance.user.address
}

resource "aws_ssm_parameter" "user_db_port" {
  name  = "/terraform/rds/user-db-port"
  type  = "String"
  value = aws_db_instance.user.port
}

resource "aws_ssm_parameter" "user_db_name" {
  name  = "/terraform/rds/user-db-name"
  type  = "String"
  value = aws_db_instance.user.db_name
}

resource "aws_ssm_parameter" "user_db_username" {
  name  = "/terraform/rds/user-db-username"
  type  = "String"
  value = aws_db_instance.user.username
}

resource "aws_ssm_parameter" "user_db_password" {
  name  = "/terraform/rds/user-db-password"
  type  = "String"
  value = aws_db_instance.user.password
}