variable "namespace" {
  description = "The project namespace to use for unique resource naming"
  type        = string
}

variable "region" {
  description = "AWS region"
  type        = string
}

variable "profile" {
  description = "AWS profile to be used"
  type        = string
}

variable "environment" {
  description = "Either dev/local/production"
  type        = string
}
