#!/bin/bash

echo 'Destroying local environment...'

set -e

export AWS_PROFILE=rental-storage

function destroy_services() {
    echo "Running destroy-service.sh script...";
    (bash ./scripts/destroy-service.sh)
}

function destroy_terraform() {
    echo "Destroying terraform configurations...";
    (cd terraform && terraform destroy -auto-approve -var-file=terraform.tfvars)
}

destroy_services
destroy_terraform
