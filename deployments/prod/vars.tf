variable "region" {
  default = "eu-central-1"
}

variable "environment" {
  default = "prod"
}

variable "base_domain_name" {}

variable "lambda_name" {}

variable "lambda_version" {}

variable "default_folder" {
  default = "/images"
}

data "aws_route53_zone" "domain_zone" {
  name = var.base_domain_name
}

locals {
  lambda_name       = var.lambda_name
  domain_name       = "${local.lambda_name}.${var.base_domain_name}"
  repo_name         = "lambda-repo.${var.base_domain_name}"
  media_bucket_name = "media.${var.base_domain_name}"
}