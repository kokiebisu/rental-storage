#!/bin/bash

function setup_terraform_versions() {
    environment=$1
    echo "Setting up versions.tf for ${environment} environment..."
    (rm terraform/versions.tf && cp terraform/config/local.tf terraform/versions.tf);
}

function setup_terraform_providers() {
    environment=$1
    echo "Setting up providers.tf for ${environment} environment..."
    (rm terraform/providers.tf && cp terraform/providers/local.tf terraform/providers.tf);
}

setup_terraform_versions $1
setup_terraform_providers $1