resource "aws_security_group" "serverless" {
  description = "Security group for Serverless functions"
  vpc_id = aws_vpc.this.id
  
  ingress {
    from_port = 0
    to_port = 65535
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "SecurityGroupServerless"
  }
}