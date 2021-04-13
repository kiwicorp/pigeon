#!/usr/bin/env bash

# scripts/build.sh
#
# This script builds binaries for pigeon.
#
# By default, binaries are built for linux/amd64 - the platform and architecture
# required by AWS Lambdas. To override the default settings, explicitly declare
# `GOOS` and `GOARCH`.

set -e

function print_help() {
    cat <<EOF
scripts/build.sh - Build pigeon targets.
    Usage: scripts/build.sh <version> <targets..>
EOF
}

# Build a target.
function build() {
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
    local outfile
    outfile="$(basename ${target})_${version}_${platform}_${arch}"

    GOOS="${platform}" GOARCH="${GOARCH}" \
        go build \
            -o "${outdir}/${outfile}" \
            "${target}"

    local new_buildid
    new_buildid="$(go tool buildid ${outdir}/${outfile})"
    local old_buildid
    if [[ -f "${outdir}/${outfile}.buildid" ]]; then
        old_buildid="$(cat ${outdir}/${outfile}.buildid 2> /dev/null)"
    else
        old_buildid=""
    fi

    if [[ "${new_buildid}" != "${old_buildid}" ]]; then
        printf "%s" "${new_buildid}" > "${outdir}/${outfile}.buildid"
    fi
}

function main() {
    local version
    version="${1}"; shift
    if [[ -z "${version}" ]]; then
        printf "build: %s\n" "no version specified\n"
        print_help
        exit 1
    fi
    if [[ -z "$@" ]]; then
        printf "build: %s\n" "no targets specified\n"
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

    local package
    package="$(head -n 1 < go.mod | cut -d ' ' -f 2)"

    printf "build: %s\n" "${package} ${version} ${platform}/${arch}"

    for target in $@; do
        printf "build: %s\n" "building ${target}"

        build "${package}/cmd/${target}" "${version}" "${platform}" "${arch}"
    done
}

main $@
