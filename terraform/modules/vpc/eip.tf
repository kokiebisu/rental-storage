resource "aws_eip" "this" {
  vpc = true

  tags = {
    Name = "EIP"
  }
}