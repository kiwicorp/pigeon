#!/usr/bin/env bash

# scripts/clean.sh
#
# This script cleans build and package artifacts.

function print_help() {
    cat <<EOF
scripts/clean.sh - Clean pigeon artifacts.
    Usage: scripts/build.sh <artifact_class> <versions> <platform_arch_list>
                <targets..>

    Artifact class
        - bin: clean binary artifacts
        - zip: clean packaged binariy artifacts
        - all: clean all artifacts

    Version
        - <versions>: comma-separated list of versions to clean
        - '-': clean all versions

    Platforms
        - <platform_arch_list>: comma-separated list of platform and
          architecture combinations to clean (linux/amd64, darwin/amd64 etc.)
        - '-': clean all platform/arch combinations

    NOTE: Work in progress. Right now, it cleans everything.
EOF
}

function main() {
    for target in $(ls cmd); do
        rm -f cmd/${target}/${target}_*_*_*
        rm -f cmd/${target}/${target}_*_*_*.zip
        rm -f cmd/${target}/${target}_*_*_*.buildid
    done
}

main $@

