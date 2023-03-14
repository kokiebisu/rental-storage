#!/bin/bash

set -e

function deploy_services () {
    local packages=("image" "space" "booking" "slack" "user" "authentication" "authorizer")
    for package in "${packages[@]}"
    do
        echo "Deploying ${package^} service...";
        (cd "packages/service-${package}" && make deploy);
    done
}

function deploy_composition () {
    local package="composition"
    echo "Deploying ${package^}...";
    (cd "packages/${package}" && make deploy);
}

function deploy_appsync() {
    local package="appsync"
    echo "Deploying ${package^}...";
    (cd "${package}" && make deploy);
}

deploy_services
deploy_composition
deploy_appsync
