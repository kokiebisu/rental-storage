variable "namespace" {
  description = "The project namespace to use for unique resource naming"
  type        = string
}

variable "booking_table_name" {
  description = "Table of name for booking service db"
  type        = string
}

variable "environment" {
  description = "Either dev/local/production"
  type        = string
}

variable "listing_db_username" {
  type = string
}

variable "listing_db_password" {
  type = string
}

variable "listing_db_name" {
  type = string
}

variable "user_db_username" {
  type = string
}

variable "user_db_password" {
  type = string
}

variable "user_db_name" {
  type = string
}