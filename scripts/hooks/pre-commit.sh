#!/bin/bash

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

lint_services
lint_composition

echo "--------------------------------------------"