#!/bin/bash

set -e

# Check if environment is local
# If it is not, exit the script
if [ "$1" != "local" ]
then
    echo "Detected environment is not local..."
    exit 1
fi

function setup_terraform_config() {
    local ENVIRONMENT=$1
    echo "Setting up terraform configuration for ${ENVIRONMENT} environment..."
    (./scripts/setup-terraform-config.sh $ENVIRONMENT);
}

function initialize() {
    local ENVIRONMENT=$1
    # Check if terraform is installed or environment is not local
    if ! command -v terraform &> /dev/null || [ "$ENVIRONMENT" != "local" ]
    then
        if ! command -v terraform &> /dev/null
        then
            echo "Terraform could not be found..."
        fi
        if [ "$1" != "local" ]
        then
            echo "Detected environment is not local..."
        fi
        echo "Installing terraform..."
        (cd terraform && terraform init);
    fi
}

function deploy() {
    echo "Deploying..."
    (cd terraform && terraform apply -auto-approve -var-file=terraform.tfvars);
}

setup_terraform_configuration local
initialize local
deploy

