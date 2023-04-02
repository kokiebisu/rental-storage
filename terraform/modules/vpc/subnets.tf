# Public Subnet
resource "aws_subnet" "a" {
  vpc_id = aws_vpc.this.id
  availability_zone = "us-east-1a"
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "subnet-a"
  }
}


resource "aws_subnet" "b" {
  vpc_id = aws_vpc.this.id
  availability_zone = "us-east-1a"
  cidr_block = "10.0.2.0/24"

  tags = {
    Name = "subnet-b"
  }
}

# Private Subnet
resource "aws_subnet" "c" {
  vpc_id = aws_vpc.this.id
  availability_zone = "us-east-1c"
  cidr_block = "10.0.3.0/24"
  tags = {
    Name = "subnet-c"
  }
}

resource "aws_subnet" "d" {
  vpc_id = aws_vpc.this.id
  availability_zone = "us-east-1d"
  cidr_block = "10.0.4.0/24"
  tags = {
    Name = "subnet-d"
  }
}