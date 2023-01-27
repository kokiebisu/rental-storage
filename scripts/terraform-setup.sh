#!/bin/bash

function setup_terraform_versions() {
    local ENVIRONMENT=$1
    local FILE="terraform/versions.tf"
    if [ -f "$FILE" ]; then
        echo "versions.tf does exist"
        (rm terraform/versions.tf)
    fi
    echo "Setting up versions.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/versions.tf terraform/versions.tf);
}

function setup_terraform_providers() {
    local ENVIRONMENT=$1
    local FILE="terraform/providers.tf"
    if [ -f "$FILE" ]; then
        echo "providers.tf does exist"
        (rm terraform/providers.tf)
    fi
    echo "Setting up providers.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/providers.tf terraform/providers.tf);
}

function setup_terraform_variables() {
    local ENVIRONMENT=$1
    local FILE="terraform/variables.tf"
    if [ -f "$FILE" ]; then
        echo "variables.tf does exist"
        (rm terraform/variables.tf)
    fi
    echo "Setting up variables.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/variables.tf terraform/variables.tf);
}

function setup_terraform_backend() {
    local ENVIRONMENT=$1
    local FILE="terraform/backend.tf"
    if [ -f "$FILE" ]; then
        echo "backend.tf does exist"
        (rm terraform/backend.tf)
    fi
    echo "Setting up backend.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/backend.tf terraform/backend.tf)
}

function remove_terraform_backend() {
    local FILE="terraform/backend.tf"
    if [ -f "$FILE" ]; then
        echo "backend.tf does exist"
        echo "Removing backend.tf..."
        (rm terraform/backend.tf)
    fi
}

setup_terraform_versions $1
setup_terraform_providers $1
setup_terraform_variables $1

if [ "$1" = "local" ]
then
    echo "Setting up for local deployment"
    remove_terraform_backend $1
fi

if [ "$1" = "dev" ] || [ "$1" = "production" ]
then
    echo "Setting up for dev/production deployment..."
    setup_terraform_backend $1
fi