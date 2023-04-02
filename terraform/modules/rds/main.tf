resource "aws_db_instance" "space" {
    username = "${var.space_db_username}"
    password = "${var.space_db_password}"
    allocated_storage = 20
    engine_version = 12.11
    db_name = "${var.space_db_name}"
    instance_class = "db.t3.micro"
    identifier = "${var.environment}-${var.namespace}-space-db"
    vpc_security_group_ids = [ "${var.db_security_group_id}" ]
    db_subnet_group_name = "${var.db_subnet_group_name}"
    engine = "postgres"
    # publicly_accessible = true
    delete_automated_backups = true
    deletion_protection = false
    skip_final_snapshot = true
}

resource "aws_db_instance" "user" {
    username = "${var.user_db_username}"
    password = "${var.user_db_password}"
    allocated_storage = 20
    engine_version = 12.11
    db_name = "${var.user_db_name}"
    instance_class = "db.t3.micro"
    identifier = "${var.environment}-${var.namespace}-user-db"
    vpc_security_group_ids = [ "${var.db_security_group_id}" ]
    db_subnet_group_name = "${var.db_subnet_group_name}"
    engine = "postgres"
    # publicly_accessible = true
    delete_automated_backups = true
    deletion_protection = false
    skip_final_snapshot = true
}
