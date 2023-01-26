#!/bin/bash

function lint_services() {
    local packages=("booking" "listing" "payment" "user" "slack")
    for package in "${packages[@]}"
    do
        echo -e "Linting ${package} service...\n"
        (cd "packages/service-${package}" && make lint)
    done
}

function lint_composition() {
    echo "Linting composition service..."
}

# lint
lint_services