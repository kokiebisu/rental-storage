variable "namespace" {
  type = string
}

variable "environment" {
  type = string
}

variable "db_security_group_id" {
  type = string
}

variable "db_subnet_group_name" {
  type = string
}

variable "user_db_info" {
  type = object({
    username = string
    password = string
    db_name  = string
  })
}

variable "space_db_info" {
  type = object({
    username = string
    password = string
    db_name  = string
  })
}
