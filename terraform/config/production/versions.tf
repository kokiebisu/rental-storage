terraform {
  required_version = "1.3.9"

  backend "s3" {
    bucket = "rental-storage-terraform-production"
    key = "production/terraform.tfstate"
    region = "us-east-1"
    encrypt = true
    kms_key_id = "alias/terraform-bucket-key"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.57.0"
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