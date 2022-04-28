module "media_bucket" {
  source        = "git::git@bitbucket.org:christian_m/aws_s3_bucket_private.git?ref=v1.0"
  environment   = var.environment
  bucket_name   = local.media_bucket_name
  force_destroy = true
  objects       = {
    cm010059_jpg = {
      document_name = "images/fotos/CM010059.jpg",
      source        = "${path.module}/assets/fotos/CM010059.jpg",
      content_type  = "image/jpg",
    },
    cm010613_jpg = {
      document_name = "images/fotos/CM010613.jpg",
      source        = "${path.module}/assets/fotos/CM010613.jpg",
      content_type  = "image/jpg",
    },
    cm010749_jpg = {
      document_name = "images/fotos/CM010749.jpg",
      source        = "${path.module}/assets/fotos/CM010749.jpg",
      content_type  = "image/jpg",
    },
    cm100250_jpg = {
      document_name = "images/fotos/CM100250.jpg",
      source        = "${path.module}/assets/fotos/CM100250.jpg",
      content_type  = "image/jpg",
    },
    cm101045_jpg = {
      document_name = "images/fotos/CM101045.jpg",
      source        = "${path.module}/assets/fotos/CM101045.jpg",
      content_type  = "image/jpg",
    },
    cm101514_jpg = {
      document_name = "images/fotos/CM101514.jpg",
      source        = "${path.module}/assets/fotos/CM101514.jpg",
      content_type  = "image/jpg",
    },
    cm101533_jpg = {
      document_name = "images/fotos/CM101533.jpg",
      source        = "${path.module}/assets/fotos/CM101533.jpg",
      content_type  = "image/jpg",
    },
    cm101534_jpg = {
      document_name = "images/fotos/CM101534.jpg",
      source        = "${path.module}/assets/fotos/CM101534.jpg",
      content_type  = "image/jpg",
    },
    cm101535_jpg = {
      document_name = "images/fotos/CM101535.jpg",
      source        = "${path.module}/assets/fotos/CM101535.jpg",
      content_type  = "image/jpg",
    },
  }
}

module "lambda_function" {
  source                = "git::git@bitbucket.org:christian_m/aws_lambda_deploy_bucket.git?ref=v1.0.1"
  environment           = var.environment
  repo_bucket           = local.repo_name
  lambda_name           = local.lambda_name
  handler               = var.lambda_name
  lambda_version        = var.lambda_version
  lambda_runtime        = "go1.x"
  memory_size           = "1536"
  timeout               = "15"
  environment_variables = {
    BUCKET_NAME    = module.media_bucket.bucket
    DEFAULT_FOLDER = var.default_folder
    ENV            = var.environment
  }
}

module "http_api" {
  source                     = "git::git@bitbucket.org:christian_m/aws_lambda_api_gateway_http.git?ref=v1.2.1"
  environment                = var.environment
  lambda_name                = local.lambda_name
  lambda_function_invoke_arn = module.lambda_function.lambda_function_invoke_arn
  domain_name                = local.domain_name
  zone_id                    = data.aws_route53_zone.domain_zone.zone_id
}

module "lambda_s3_access" {
  source                     = "git::git@bitbucket.org:christian_m/aws_s3_lambda_access_policy.git?ref=v1.0"
  environment                = var.environment
  bucket_arn                 = module.media_bucket.bucket_arn
  lambda_name                = local.lambda_name
  lambda_execution_role_name = module.lambda_function.lambda_execution_role_name
}