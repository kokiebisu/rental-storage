#!/bin/bash

function setup_terraform_versions() {
    environment=$1
    echo "Setting up versions.tf for ${environment} environment..."
    (rm terraform/versions.tf && cp terraform/config/${environment}.tf terraform/versions.tf);
}

function setup_terraform_providers() {
    environment=$1
    echo "Setting up providers.tf for ${environment} environment..."
    (rm terraform/providers.tf && cp terraform/providers/${environment}.tf terraform/providers.tf);
}

function setup_terraform_variables() {
    environment=$1
    echo "Setting up variables.tf for ${environment} environment..."
    (rm terraform/variables.tf && cp terraform/variables/${environment}.tf terraform/variables.tf);
}

setup_terraform_versions $1
setup_terraform_providers $1
setup_terraform_variables $1