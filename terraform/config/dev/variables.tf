variable "namespace" {
  description = "The project namespace to use for unique resource naming"
  type        = string
}

variable "region" {
  description = "AWS region"
  type        = string
}

variable "environment" {
  description = "Either dev/local/production"
  type        = string
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

variable "booking_db_info" {
  type = object({
    table_name = string
  })
}

variable "chat_db_info" {
  type = object({
    table_name = string
  })
}
