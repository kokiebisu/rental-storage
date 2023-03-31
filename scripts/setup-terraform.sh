#!/bin/bash

set -e

function setup_terraform_versions() {
    local ENVIRONMENT=$1
    local FILE="terraform/versions.tf"
    if [ -f "$FILE" ]; then
        (rm terraform/versions.tf)
    fi
    echo "Setting up versions.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/versions.tf terraform/versions.tf);
}

function setup_terraform_providers() {
    local ENVIRONMENT=$1
    local FILE="terraform/providers.tf"
    if [ -f "$FILE" ]; then
        (rm terraform/providers.tf)
    fi
    echo "Setting up providers.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/providers.tf terraform/providers.tf);
}

function setup_terraform_variables() {
    local ENVIRONMENT=$1
    local FILE="terraform/variables.tf"
    if [ -f "$FILE" ]; then
        (rm terraform/variables.tf)
    fi
    echo "Setting up variables.tf for ${ENVIRONMENT} environment..."
    (cp terraform/config/${ENVIRONMENT}/variables.tf terraform/variables.tf);
}

function deploy() {
    echo "Deploying..."
    # Check if terraform is installed
    if ! command -v terraform &> /dev/null
    then
        echo "Terraform could not be found"
        # Install terraform
        echo "Installing terraform..."
        (cd terraform && terraform init)
    fi
    (cd terraform && terraform apply -auto-approve -var-file=terraform.tfvars)
}

setup_terraform_versions $1
setup_terraform_providers $1
setup_terraform_variables $1
deploy

if [ "$1" = "local" ]
then
    echo "Setting up for local deployment"
fi

if [ "$1" = "dev" ] || [ "$1" = "production" ]
then
    echo "Setting up for dev/production deployment..."
fi