#!/bin/bash

echo 'Setting up local environment...'

set -e

export AWS_PROFILE=rental-storage

function check_docker_running() {
    # if docker is not running exit
    if ! docker ps > /dev/null 2>&1; then
        echo "Docker is not running. Please start docker and try again."
        exit 1
    fi
}

function vscode_prepare() {
    echo "Setting up vscode configurations...";
    (mkdir -p .vscode && cp config/settings.json .vscode/settings.json);
}

function install_dependencies() {
    echo "Installing dependencies...";
    (bash ./scripts/setup-deps.sh)
}

function setup_terraform() {
    echo "Setting up terraform configurations...";
    (bash ./scripts/setup-terraform.sh $1)
}

function setup_serverless() {
    echo "Setting up serverless configurations...";
    (bash ./scripts/setup-service.sh)
}

check_docker_running
vscode_prepare
install_dependencies
setup_terraform local
setup_serverless
