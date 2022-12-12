resource "aws_ssm_parameter" "listing_db_host" {
    name = "/terraform/rds/listing-db-host"
    type = "String"
    value = aws_db_instance.listing.address
}

resource "aws_ssm_parameter" "listing_db_port" {
    name = "/terraform/rds/listing-db-port"
    type = "String"
    value = aws_db_instance.listing.port
}

resource "aws_ssm_parameter" "listing_db_name" {
    name = "/terraform/rds/listing-db-name"
    type = "String"
    value = aws_db_instance.listing.name
}

resource "aws_ssm_parameter" "listing_db_username" {
    name = "/terraform/rds/listing-db-username"
    type = "String"
    value = aws_db_instance.listing.username
}

resource "aws_ssm_parameter" "listing_db_password" {
    name = "/terraform/rds/listing-db-password"
    type = "String"
    value = aws_db_instance.listing.password
}

resource "aws_ssm_parameter" "user_db_host" {
    name = "/terraform/rds/user-db-host"
    type = "String"
    value = aws_db_instance.user.address
}

resource "aws_ssm_parameter" "user_db_port" {
    name = "/terraform/rds/user-db-port"
    type = "String"
    value = aws_db_instance.user.port
}

resource "aws_ssm_parameter" "user_db_name" {
    name = "/terraform/rds/user-db-name"
    type = "String"
    value = aws_db_instance.user.name
}

resource "aws_ssm_parameter" "user_db_username" {
    name = "/terraform/rds/user-db-username"
    type = "String"
    value = aws_db_instance.user.username
}

resource "aws_ssm_parameter" "user_db_password" {
    name = "/terraform/rds/user-db-password"
    type = "String"
    value = aws_db_instance.user.password
}