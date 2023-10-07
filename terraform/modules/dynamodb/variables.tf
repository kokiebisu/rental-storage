variable "namespace" {
  type = string
}

variable "environment" {
  type = string
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
