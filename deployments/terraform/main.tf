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

provider "cloudflare" {
  api_token = var.cf_api_token
}

module "gh_content_fn" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "1.45.0"

  function_name = "content_github_releases"
  description   = "GitHub Releases content function."
  handler       = "cmd/content_github_releases/content_github_releases_dev_linux_amd64"
  runtime       = "go1.x"

  create_package         = false
  local_existing_package = "../../cmd/content_github_releases/content_github_releases_dev_linux_amd64.zip"

  allowed_triggers = { for key, value in module.gh_content_events.eventbridge_rule_arns: key => {
      service = "events"
      arn     = value
    }
  }
  # fixme 24/04/2021: only for unpublished versions
  create_current_version_allowed_triggers = false

  attach_policy_statements = true
  policy_statements = {
    dynamodb = {
      effect    = "Allow",
      actions   = [
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
        "dynamodb:DeleteItem",
      ],
      resources = [module.content_table.this_dynamodb_table_arn]
    }
    dynamodb_index = {
      effect    = "Allow",
      actions   = [
        "dynamodb:Query",
        "dynamodb:Scan",
      ],
      resources = ["${module.content_table.this_dynamodb_table_arn}/index/*"]
    }
    ses = {
      effect  = "Allow"
      actions = [
        "ses:SendEmail",
        "ses:SendRawEmail",
      ]
      resources = ["*"]
    }
  }

  tags = local.tags
}

module "gh_content_events" {
  source  = "terraform-aws-modules/eventbridge/aws"
  version = "1.1.0"

  create_bus = false

  rules   = local.event_rules
  targets = local.event_targets

  tags = local.tags
}

module "content_table" {
  source = "terraform-aws-modules/dynamodb-table/aws"

  name     = "content-github-releases"
  # fixme 24/04/2021: hardcoded hash key
  hash_key = "Urn"

  attributes = [
    {
      # fixme 24/04/2021: hardcoded hash key
      name = "Urn" # urn:pigeon.selftech.io:content:github_releases:/<repo_owner>/<repo_name>
      type = "S"
    }
  ]

  tags = local.tags
}

locals {
  tags = {
    PigeonVersion     = "dev"
    PigeonReleaseName = "gh-content-demo-mvp"
  }

  event_rule_names = [for gh_content in var.gh_content : "event_rule_${gh_content.repo_owner}_${gh_content.repo_name}"]
  event_target_names = [for gh_content in var.gh_content : "event_target_${gh_content.repo_owner}_${gh_content.repo_name}"]

  event_rules = { for index, gh_content in var.gh_content : local.event_rule_names[index] => {
    description         = "GitHub content - ${gh_content.repo_owner}/${gh_content.repo_name}."
    schedule_expression = "cron(0 0 ? * * *)"
  } }

  event_targets = { for index, gh_content in var.gh_content : local.event_rule_names[index] => [{
    name  = local.event_target_names[index]
    arn   = module.gh_content_fn.this_lambda_function_arn
    input = jsonencode({
      access_token = var.gh_access_token
      repo_owner   = gh_content.repo_owner
      repo_name    = gh_content.repo_name
    })
  }] }
}

module "pigeon_ses" {
  source  = "selftechio/ses-cf-domain/aws"
  version = "0.1.0"

  zone   = var.ses_zone
  domain = var.ses_domain
}
