# variables.tf

variable "aws_creds_file" {
  type        = string
  description = "Credentials file used for authenticating with the AWS provider."
}

variable "aws_region" {
  type        = string
  description = "AWS region to deploy to."
}

variable "gh_access_token" {
  type        = string
  description = "GitHub access token used for making API calls to GitHub."
}

variable "gh_content" {
  type = list(object({
    repo_owner = string
    repo_name  = string
  }))
  description = "A list of content parameters used for building events for GitHub content."
}
