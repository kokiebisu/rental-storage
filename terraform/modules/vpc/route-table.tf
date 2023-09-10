# Public
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }

  tags = {
    Name = "route-table-public"
  }
}

resource "aws_route_table_association" "public_a" {
  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.a.id
}

resource "aws_route_table_association" "public_c" {
  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.c.id
}

# Private
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.this.id


  tags = {
    Name = "route-table-private"
  }
}

resource "aws_route_table_association" "private_b" {
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.b.id
}

resource "aws_route_table_association" "private_d" {
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.d.id
}
