resource "aws_key_pair" "bastion" {
  key_name   = "my_keypair"
  public_key = file("~/.ssh/id_rsa_rental_storage_ec2.pub")
}

resource "aws_launch_template" "bastion" {
  image_id      = "ami-053b0d53c279acc90"
  instance_type = "t2.micro"
  key_name = aws_key_pair.bastion.key_name
  vpc_security_group_ids = [var.security_group_id]
}

resource "aws_autoscaling_group" "bar" {
  vpc_zone_identifier = [ var.primary_public_subnet_id ]
  # availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d"]
  desired_capacity   = 1
  max_size           = 1
  min_size           = 1

  launch_template {
    id      = aws_launch_template.bastion.id
    version = "$Latest"
  }
}
