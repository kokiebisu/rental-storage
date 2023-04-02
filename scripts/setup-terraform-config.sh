#!/bin/bash

set -e


# Check if environment is specified
# If it is not, exit the script
if [ -z "$1" ]
then
    echo "Environment is not specified. Specify environment './scripts/setup-terraform-config.sh <environment>'."
    exit 1
fi

# if the environment is not local, prompt the user if they are sure they want to continue
if [ "$1" != "local" ]
then
    echo "You specified an environment that is not local."
    read -p "Are you sure you want to continue? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]
    then
        exit 1
    fi
fi

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

setup_terraform_versions $1
setup_terraform_providers $1
setup_terraform_variables $1
