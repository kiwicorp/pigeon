#!/usr/bin/env bash

# scripts/package.sh
#
# This script packages binaries for pigeon.
#
# AWS Lambda requires binaries to be built for linux/amd64. Therefore, packaging
# will only work for that platform and architecture combination.

function print_help() {
    cat <<EOF
scripts/package.sh - Package pigeon targets
    Usage: scripts/package.sh <version> <targets..>
EOF
}

# Package a target.
function package() {
    local target
    target="${1}"; shift
    local version
    version="${1}"; shift
    local platform
    platform="${1}"; shift
    local arch
    arch="${1}"; shift

    local outdir
    outdir="cmd/$(basename ${target})"
    local filename
    filename="$(basename ${target})_${version}_${platform}_${arch}"

    zip "${outdir}/${filename}.zip" "${outdir}/${filename}"
}

function check_if_already_up_to_date() {
    local target
    target="${1}"; shift
    local version
    version="${1}"; shift
    local platform
    platform="${1}"; shift
    local arch
    arch="${1}"; shift

    local outdir
    outdir="cmd/$(basename ${target})"
    local filename
    filename="$(basename ${target})_${version}_${platform}_${arch}"

    local buildid_mtime
    buildid_mtime="$(date -r ${outdir}/${filename} +%s)"
    local zip_mtime
    if [[ ! -f "${outdir}/${filename}.zip" ]]; then
        zip_mtime="0"
    else
        zip_mtime="$(date -r ${outdir}/${filename}.zip +%s)"
    fi

    if [[ ${buildid_mtime} -lt 1 || ${zip_mtime} -gt 0 ]]; then
        printf "%s" "true"
    else
        printf "%s" "false"
    fi
}

function main() {
    local version
    version="${1}"; shift
    if [[ -z "${version}" ]]; then
        printf "scripts/package.sh: %s\n" "no version specified\n"
        print_help
        exit 1
    fi
    if [[ -z "$@" ]]; then
        printf "scripts/package.sh: %s\n" "no targets specified\n"
        print_help
        exit 1
    fi

    if [[ -z "${GOOS}" ]]; then
        export GOOS="linux"
    fi
    local platform
    platform="$(go env GOOS)"
    if [[ -z "${GOARCH}" ]]; then
        export GOARCH="amd64"
    fi
    local arch
    arch="$(go env GOARCH)"

    if [[ "${platform}/${arch}" != "linux/amd64" ]]; then
        printf "scripts/package.sh: %s, %s\n\n" \
            "expected linux/amd64" \
            "refusing to package for ${platform}/${arch}"
        print_help
        exit 1
    fi

    local package
    package="$(head -n 1 < go.mod | cut -d ' ' -f 2)"

    printf "scripts/package.sh: %s\n" "${package} ${version} ${platform}/${arch}"

    for target in $@; do
        printf "scripts/package.sh: %s\n" "packaging ${target}"

        local utd
        utd="$(check_if_already_up_to_date ${package}/cmd/${target} ${version} ${platform} ${arch})"
        if [[ "${utd}" == "true" ]]; then
            printf "scripts/package.sh: %s\n" "already up to date"
            continue
        fi

        package "${package}/cmd/${target}" "${version}" "${platform}" "${arch}"
    done
}

main $@
