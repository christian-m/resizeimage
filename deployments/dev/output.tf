output "internal_http_api_url" {
  value = module.http_api.api_url
}

output "external_http_api_url" {
  value = "https://${local.domain_name}/"
}