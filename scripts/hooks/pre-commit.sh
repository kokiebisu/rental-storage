#!/bin/bash

set -e

function lint_platforms() {
    local packages=("web")
    for package in "${packages[@]}"
    do
        echo -e "\nLinting ${package}..."
        (cd "${package}" && make lint)
    done
}

function lint_services() {
    local packages=("authentication" "authorizer" "booking" "listing" "payment" "user" "slack")
    for package in "${packages[@]}"
    do
        echo -e "\nLinting ${package} service..."
        (cd "packages/service-${package}" && make lint)
    done
}

function lint_composition() {
    local package="composition"
    echo -e "\nLinting composition service..."
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
