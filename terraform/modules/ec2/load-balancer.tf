resource "aws_lb" "service_search" {
  name               = "service-search-alb"
  internal           = false
  load_balancer_type = "application"
  subnets            = [var.primary_public_subnet_id, var.secondary_public_subnet_id]
  security_groups    = [var.alb_security_group_id]
}

resource "aws_lb_target_group" "service_search" {
  name     = "service-search-target-group"
  port     = 8080
  protocol = "HTTP"
  vpc_id   = var.vpc_id
}

resource "aws_lb_listener" "service_search" {
  load_balancer_arn = aws_lb.service_search.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.service_search.arn
    type             = "forward"
  }
}
