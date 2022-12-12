resource "aws_subnet" "a" {
  vpc_id = aws_vpc.default.id
  availability_zone = "us-east-1a"
  cidr_block = "10.0.0.0/24"

  tags = {
    Name = "SubnetA"
  }
}

resource "aws_subnet" "b" {
  vpc_id = aws_vpc.default.id
  availability_zone = "us-east-1b"
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "SubnetB"
  }
}

resource "aws_subnet" "c" {
  vpc_id = aws_vpc.default.id
  availability_zone = "us-east-1c"
  cidr_block = "10.0.2.0/24"

  tags = {
    Name = "SubnetC"
  }
}