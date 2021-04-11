#!/usr/bin/env bash

# tf_wrapper.sh
#
# This script is a wrapper around terraform for common operations.
#
# WORK IN PROGRESS!
#
# Supported operations:
# - apply
# - auto-apply
# - destroy
# - plan


set -e

DEPLOYMENT_ENV="${1}"; shift
stat "deployments/terraform/env/${DEPLOYMENT_ENV}.tfvars" > /dev/null 2>&1 \
    || {
        printf "%s\n" "No such profile: ${DEPLOYMENT_ENV}. Creating env/${DEPLOYMENT_ENV}.tfvars."
        cp \
            deployments/terraform/env/sample.tfvars \
            "deployments/terraform/env/${DEPLOYMENT_ENV}.tfvars"
        terraform -chdir=deployments/terraform \
            workspace new "${DEPLOYMENT_ENV}"
        exit 1
    }

terraform -chdir=deployments/terraform workspace select "${DEPLOYMENT_ENV}"

terraform -chdir=deployments/terraform \
    apply --var-file="env/${DEPLOYMENT_ENV}.tfvars"
