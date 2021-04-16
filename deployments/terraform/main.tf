# main.tf

provider "aws" {
  shared_credentials_file = var.aws_creds_file
  region                  = var.aws_region

  # Make it faster by skipping something
  skip_get_ec2_platforms      = true
  skip_metadata_api_check     = true
  skip_region_validation      = true
  skip_credentials_validation = true
  skip_requesting_account_id  = true
}

module "gh_content_fn" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "1.45.0"

  function_name = "lambda-function-test"
  description   = "A test lambda function."
  handler       = "cmd/lambda/lambda-linux-amd64"
  runtime       = "go1.x"

  create_package         = false
  local_existing_package = "../../cmd/lambda/lambda_dev_linux_amd64.zip"
}

module "gh_content_events" {
  source  = "terraform-aws-modules/eventbridge/aws"
  version = "1.1.0"

  create_bus = false

  rules = local.event_rules

  targets = local.event_targets
}

locals {
  event_rule_names = [for gh_content in var.gh_content : "event_${gh_content.repo_owner}_${gh_content.repo_name}"]

  event_rules = { for index, gh_content in var.gh_content : local.event_rule_names[index] => {
    description         = "GitHub content - ${gh_content.repo_owner}/${gh_content.repo_name}."
    schedule_expression = "cron(0 0 ? * * *)"
  } }

  event_targets = { for index, gh_content in var.gh_content : local.event_rule_names[index] => [ {
    name     = "Execute lambda for ${gh_content.repo_owner}/${gh_content.repo_name}."
    arn      = module.gh_content_fn.this_lambda_function_arn
    constant = jsonencode({
      access_token = var.gh_access_token
      repo_owner   = gh_content.repo_owner
      repo_name    = gh_content.repo_name
    })
  } ] }
}
