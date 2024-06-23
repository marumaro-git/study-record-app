locals {
  region      = "ap-northeast-1"
  aws_profile = "default"
}

provider "aws" {
  region  = local.region
  profile = local.aws_profile
}

terraform {
  required_version = ">= 1.8.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.40.0"
    }
  }
}
