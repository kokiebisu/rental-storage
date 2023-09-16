variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "flow_logs_role" {
  type = string
}

variable "account_id" {
  type = string
}

variable "public_ec2_id" {
  type = string
}
