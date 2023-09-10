resource "aws_ssm_parameter" "lb_dns_name" {
  name  = "/terraform/ec2/lb-dns-name"
  type  = "String"
  value = aws_lb.service_search.dns_name
}
