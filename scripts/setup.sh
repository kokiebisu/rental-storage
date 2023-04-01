#!/bin/bash

echo 'Setting up local environment...'

set -e

export AWS_PROFILE=rental-storage

function check_requirements() {
    echo "Checking requirements...";
    # check if the script is called from the root directory
    if [ ! -f "scripts/setup.sh" ]; then
        echo "Please run this script from the root directory of the project.";
        exit 1;
    fi

    # check if docker is running
    if ! docker ps > /dev/null 2>&1; then
        echo "Docker is not running. Please start docker and try again.";
        exit 1;
    fi
}

function setup_vscode() {
    echo "Setting up vscode configurations...";
    (mkdir -p .vscode && cp config/settings.json .vscode/settings.json);
}

function setup_deps() {
    echo "Installing dependencies...";
    (bash ./scripts/setup-deps.sh);
}

function setup_terraform() {
    echo "Setting up terraform configurations...";
    (bash ./scripts/setup-terraform.sh $1);
}

function setup_serverless() {
    echo "Setting up serverless configurations...";
    (bash ./scripts/setup-service.sh);
}

check_requirements
setup_vscode
setup_deps
setup_terraform local
setup_serverless
