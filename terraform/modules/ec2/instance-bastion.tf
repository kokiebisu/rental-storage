resource "aws_instance" "bastion_public" {
  ami                         = "ami-007855ac798b5175e"
  instance_type               = "t2.micro"
  key_name                    = aws_key_pair.my_keypair.key_name
  vpc_security_group_ids      = [var.public_ec2_security_group_id]
  subnet_id                   = var.primary_public_subnet_id
  associate_public_ip_address = true

  connection {
    type        = "ssh"
    user        = "ubuntu"                                 # The default username for Ubuntu, replace with your desired username if you're using a custom AMI
    private_key = file("~/.ssh/id_rsa_rental_storage_ec2") # Replace with the path to your private key file
    host        = self.public_ip
  }

  tags = {
    Name = "public-bastion-host"
  }
}

# resource "aws_instance" "bastion_private" {
#   ami                         = "ami-007855ac798b5175e"
#   instance_type               = "t2.micro"
#   key_name                    = aws_key_pair.my_keypair.key_name
#   vpc_security_group_ids      = [var.private_ec2_security_group_id]
#   subnet_id                   = var.primary_private_subnet_id

#   connection {
#     type        = "ssh"
#     user        = "ubuntu"                                 # The default username for Ubuntu, replace with your desired username if you're using a custom AMI
#     private_key = file("~/.ssh/id_rsa_rental_storage_ec2") # Replace with the path to your private key file
#     host        = self.private_ip
#   }

#   tags = {
#     Name = "private-bastion-host"
#   }
# }

resource "aws_eip" "bastion_public" {
  instance = aws_instance.bastion_public.id

  tags = {
    Name = "public-bastion-host"
  }
}

# generates a json with the rds endpoints
# this will be used to set up an ssh tunnel via bastion
# to connect from local
resource "local_file" "bastion" {
  filename = "json/ec2_bastion.json"
  content = jsonencode({
    endpoint = "${aws_eip.bastion_public.public_ip}"
  })
}
