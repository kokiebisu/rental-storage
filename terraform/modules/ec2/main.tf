resource "aws_key_pair" "bastion" {
  key_name   = "my_keypair"
  public_key = file("~/.ssh/id_rsa_rental_storage_ec2.pub")
}

resource "aws_launch_template" "bastion" {
  image_id               = "ami-053b0d53c279acc90"
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.bastion.key_name
  vpc_security_group_ids = [var.security_group_id]
}

resource "aws_autoscaling_group" "bar" {
  vpc_zone_identifier = [var.primary_public_subnet_id]

  desired_capacity = 1
  max_size         = 1
  min_size         = 1

  mixed_instances_policy {
    launch_template {
      # Use a Launch Template to define the base launch configuration
      # with the required instance type and AMI ID.
      launch_template_specification {
        launch_template_id = aws_launch_template.bastion.id
        version            = "$Latest"
      }

      override {
        instance_type = "t2.micro"
      }
    }

    instances_distribution {
      spot_allocation_strategy                 = "lowest-price"
      on_demand_base_capacity                  = 1
      on_demand_percentage_above_base_capacity = 0
    }
  }
}
