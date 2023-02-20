terraform {
  required_version = ">= 0.15"

  backend "s3" {
    bucket = "rental-storage-terraform"
    key = "dev/terraform.tfstate"
    region = "us-east-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.28"
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