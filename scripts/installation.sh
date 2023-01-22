#!/bin/bash

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

function install_dependencies_appsync() {
    local package="appsync"
    echo "Installing ${package^} service dependencies...";
    (cd appsync && pnpm install);
}

function install_dependencies_composition() {
    local package="composition"
    name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
    echo "Installing ${name} dependencies...";
    (cd packages/composition && pnpm install);
}

install_dependencies_root
install_dependencies_services
install_dependencies_appsync
install_dependencies_composition