resource "local_file" "aws_db_instance" {
  filename = "json/rds.json"
  content = jsonencode({
    space = "${aws_db_instance.space.endpoint}"
    user  = "${aws_db_instance.user.endpoint}"
  })
}
