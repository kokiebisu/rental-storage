terraform {
  required_version = ">= 1.0.0, < 2.0.0"

  backend "s3" {
    bucket = "rental-storage-terraform"
    key    = "dev/terraform.tfstate"
    region = "us-east-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.16.1"
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