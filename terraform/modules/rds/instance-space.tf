resource "aws_db_instance" "space" {
  username                 = var.space_db_info.username
  password                 = var.space_db_info.password
  allocated_storage        = 20
  engine_version           = 12.14
  db_name                  = var.space_db_info.db_name
  instance_class           = "db.t3.micro"
  identifier               = "${var.environment}-${var.namespace}-space-db"
  vpc_security_group_ids   = ["${var.db_security_group_id}"]
  db_subnet_group_name     = var.db_subnet_group_name
  engine                   = "postgres"
  delete_automated_backups = true
  deletion_protection      = false
  skip_final_snapshot      = true
}
