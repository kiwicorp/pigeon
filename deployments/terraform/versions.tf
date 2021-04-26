# versions.tf

terraform {
  required_providers {
    aws        = "~> 3"
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 2"
    }
  }
  required_version = ">= 0.15"
}
