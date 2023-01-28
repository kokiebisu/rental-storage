#!/bin/bash

set -e

function install_dependencies_root() {
    echo "Installing root dependencies...";
    (pnpm install);
}

function install_dependencies_services() {
    local packages=("image" "listing" "booking" "slack" "user" "authentication" "authorizer")
    
    for package in "${packages[@]}"
    do
        echo "Installing ${service^} service dependencies...";
        (cd "packages/service-${service}" && pnpm install);
    done
}

function install_dependencies_composition() {
    local package="composition"
    echo "Installing ${package^} service dependencies...";
    (cd "packages/${package}" && pnpm install);
}

function install_dependencies_appsync() {
    local package="appsync"
    echo "Installing ${package^} service dependencies...";
    (cd appsync && pnpm install);
}

install_dependencies_root
install_dependencies_services
install_dependencies_composition
install_dependencies_appsync