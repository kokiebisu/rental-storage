#!/bin/bash

set -e

# Check if environment is provided
# If it is not, exit the script
if [ -z "$1" ]
then
    echo "Environment is not specified. Specify environment './scripts/setup-terraform.sh <environment>'."
    exit 1
fi

# if the environment is not local, prompt the user if they are sure they want to continue
if [ "$1" != "local" ]
then
    echo "You specified an environment that is not local. "
    read -p "Are you sure you want to continue? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]
    then
        exit 1
    fi
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

setup_terraform_config local
initialize local
deploy

