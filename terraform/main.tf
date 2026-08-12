terraform {
  required_version = ">= 1.6"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  default = "eu-west-2" # London
}

resource "aws_cloudwatch_log_group" "app" {
  name              = "/aws/ec2/alerting-app"
  retention_in_days = 7
}

output "log_group_name" {
  value = aws_cloudwatch_log_group.app.name
}