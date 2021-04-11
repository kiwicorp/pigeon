# variables.tf

variable "aws_creds_file" {
  type        = string
  description = "Credentials file used for authenticating with the AWS provider."
}

variable "aws_region" {
  type        = string
  description = "AWS region to deploy to."
}
