#!/usr/bin/env bash

# scripts/deploy.sh
#
# This script deploys the application to AWS.

function print_help() {
    cat <<EOF
scripts/deploy.sh - Deploy pigeon.
    Usage: scripts/deploy.sh <version>
EOF
}

function deploy() {
    local env_profile
    env_profile="${1}"; shift

    terraform \
        -chdir=deployments/terraform \
        workspace \
        select \
        "${env_profile}"

    terraform \
        -chdir=deployments/terraform \
        apply \
        --auto-approve \
        -input=true \
        --var-file="env/${env_profile}.tfvars"
}

function main() {
    local version
    version="${1}"; shift
    if [[ -z "${version}" ]]; then
        printf "deploy: %s\n" "no version specified\n"
        print_help
        exit 1
    fi

    printf "deploy: %s\n" "${version}"

    deploy "${version}"
}

main $@
