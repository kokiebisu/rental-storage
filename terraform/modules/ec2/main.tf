# use this module to create an ec2 instance for aws resource validation
resource "aws_instance" "bastion" {
  ami           = "ami-007855ac798b5175e"
  instance_type = "t2.micro"
  key_name = aws_key_pair.my_keypair.key_name
  vpc_security_group_ids = [var.security_group_id]
  subnet_id = var.primary_public_subnet_id
  associate_public_ip_address = true

  connection {
    type        = "ssh"
    user        = "ubuntu" # The default username for Ubuntu, replace with your desired username if you're using a custom AMI
    private_key = file("~/.ssh/id_rsa_rental_storage_ec2") # Replace with the path to your private key file
    host        = self.public_ip
  }

  tags = {
      Name = "bastion-host"
  }
}

resource "aws_key_pair" "my_keypair" {
  key_name   = "my_keypair"
  public_key = file("~/.ssh/id_rsa_rental_storage_ec2.pub")
}
