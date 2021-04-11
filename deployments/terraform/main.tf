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

module "lambda_function_test" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "1.45.0"

  function_name = "lambda-function-test"
  description   = "A test lambda function."
  handler       = "cmd/lambda/lambda-linux-amd64"
  runtime       = "go1.x"

  create_package         = false
  local_existing_package = "../../cmd/lambda/lambda.zip"
}

# module "eventbridge_test" {
#   source  = "terraform-aws-modules/eventbridge/aws"
#   version = "1.1.0"

#   rules = {
#     schedule_lambda = {
#       description   = "Schedule lambda execution.",
#       schedule_expression = "cron(* * * * * *)"
#     }
#   }

#   targets = {
#     schedule_lambda = [
#       {
#         arn  = module.lambda_function_test.this_lambda_function_invoke_arn
#         name = "Execute lambda function"
#       }
#     ]
#   }
# }
