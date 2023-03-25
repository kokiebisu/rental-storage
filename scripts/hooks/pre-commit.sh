#!/bin/bash

set -e

function lint_platforms() {
    local packages=("web")
    for package in "${packages[@]}"
    do
        (cd "${package}" && make lint)
    done
}

function lint_services() {
    local packages=("authentication" "authorizer" "booking" "space" "user" "slack")
    for package in "${packages[@]}"
    do
        (cd "packages/service-${package}" && make lint)
    done
}

function lint_composition() {
    local package="composition"
    (cd "packages/${package}" && make lint)
}

if [[ -z $(git diff --cached --exit-code) ]]; then
    echo "No commits to be made. Please stage some files...";
    exit 1
else
    echo "Commits to be made..."
    lint_services
    lint_composition
    lint_platforms
    echo "--------------------------------------------"
fi
