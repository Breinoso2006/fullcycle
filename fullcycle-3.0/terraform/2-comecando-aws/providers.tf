terraform {
  required_version = ">=1.14.1"
  required_providers {
    aws = ">=6.24.0"
    local = ">= 2.6.1"
  }
}

provider "aws" {
  region = "sa-east-1"
}