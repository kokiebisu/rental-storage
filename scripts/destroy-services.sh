#!/bin/bash

set -e

function remove_lambda_services() {
    local packages=("chat" "image" "space" "booking" "slack" "user" "authentication" "authorizer")
    for package in "${packages[@]}"
    do
        echo "Removing ${package^} service...";
        (cd "packages/service-${package}" && pnpm run remove);
    done
}


function remove_composition() {
    local package="composition"
    echo "Removing ${package^}...";
    (cd "packages/${package}" && pnpm run remove);
}

function remove_appsync() {
    local package="appsync"
    echo "Removing ${package^}...";
    (cd "${package}" && pnpm run remove);
}

remove_lambda_services
remove_composition
remove_appsync
