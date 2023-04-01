#!/bin/bash

set -e

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

setup_terraform_configuration $1
initialize $1
deploy

if [ "$1" = "local" ]
then
    echo "Setting up for local deployment"
fi

if [ "$1" = "dev" ] || [ "$1" = "production" ]
then
    echo "Setting up for dev/production deployment..."
fi