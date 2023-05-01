#!/bin/bash

set -e

function deploy_services () {
    local packages=("chat" "image" "space" "booking" "slack" "user" "authentication" "authorizer")
    for package in "${packages[@]}"
    do
        echo "Deploying ${package^} service...";
        (cd "packages/service-${package}" && pnpm run deploy);
    done
}

function deploy_composition () {
    local package="composition"
    echo "Deploying ${package^}...";
    (cd "packages/${package}" && pnpm run deploy);
}

function deploy_appsync() {
    local package="appsync"
    echo "Deploying ${package^}...";
    (cd "${package}" && pnpm run deploy);
}

deploy_services
deploy_composition
deploy_appsync
