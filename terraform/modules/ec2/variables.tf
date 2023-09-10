variable "namespace" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "public_ec2_security_group_id" {
  type = string
}

variable "private_ec2_security_group_id" {
  type = string
}

variable "alb_security_group_id" {
  type = string
}

variable "primary_public_subnet_id" {
  type = string
}

variable "primary_private_subnet_id" {
  type = string
}

variable "secondary_public_subnet_id" {
  type = string
}
