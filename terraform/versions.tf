terraform {
  required_version = ">= 0.15"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.61.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }

    cloudinit = {
      source  = "hashicorp/cloudinit"
      version = "~> 2.1"
    }
  }
}