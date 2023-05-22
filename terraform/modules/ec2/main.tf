# how you can configure RDS to allow access to PGAdmin while keeping the RDS instance private:

# Set up an EC2 instance:
# First, create an Amazon EC2 instance in a public subnet with a security group that allows SSH traffic from your IP address or range of IP addresses. You can use this instance as a bastion host to securely forward traffic to your RDS instance. The bastion host acts as an intermediary server that enables secure communication between your local machine and the RDS instance.

# Modify the Security Group Rules for RDS and EC2:
# Next, modify the security group rules for both your EC2 instance and RDS instance to allow inbound traffic on the PostgreSQL port (default: 5432) from your bastion host only. This ensures that no other traffic can access your RDS instance directly.

# Connect to the EC2 instance via SSH:
# Connect to your EC2 instance using SSH and forward a local port to your RDS instance. For example, you could forward the local port 5432 to the RDS instance’s port 5432 using the following command:

# ssh -i <path-to-key-file> -N -L localhost:5432:<your-rds-endpoint>:5432 ec2-user@<ec2-instance-public-ip>

# Replace <path-to-key-file> with the path to your private key file, <your-rds-endpoint> with your RDS instance endpoint, and <ec2-instance-public-ip> with the public IP address of your EC2 instance. This command forwards all traffic on localhost:5432 to your RDS instance’s PostgreSQL port.

# Access the RDS instance via PGAdmin:
# Now you can launch PGAdmin and connect to your RDS database through the forwarded port using the credentials provided by AWS. Make sure to specify localhost as the server hostname and 5432 as the port number.
# In summary, configuring RDS to allow access to PGAdmin requires setting up an EC2 instance as a bastion host to forward traffic from your local machine to the RDS instance. This approach adds an additional layer of security by ensuring that only authorized traffic can reach your RDS instance, helping to protect your database from unauthorized access.

# use this module to create an ec2 instance for aws resource validation
# resource "aws_instance" "bastion" {
#   ami           = "ami-007855ac798b5175e"
#   instance_type = "t2.micro"
#   key_name = aws_key_pair.my_keypair.key_name
#   vpc_security_group_ids = [var.security_group_id]
#   subnet_id = var.primary_public_subnet_id
#   associate_public_ip_address = true

#   connection {
#     type        = "ssh"
#     user        = "ubuntu" # The default username for Ubuntu, replace with your desired username if you're using a custom AMI
#     private_key = file("~/.ssh/id_rsa_rental_storage_ec2") # Replace with the path to your private key file
#     host        = self.public_ip
#   }

#   tags = {
#       Name = "bastion-host"
#   }
# }

# resource "aws_key_pair" "my_keypair" {
#   key_name   = "my_keypair"
#   public_key = file("~/.ssh/id_rsa_rental_storage_ec2.pub")
# }
