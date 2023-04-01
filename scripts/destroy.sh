#!/bin/bash

echo 'Destroying local environment...'

set -e

export AWS_PROFILE=rental-storage

function destroy_terraform() {
    echo "Destroying terraform configurations...";
    (cd terraform && terraform destroy -auto-approve)
}

function destroy_serverless() {
    echo "Destroying serverless configurations...";
    (bash ./scripts/destroy-service.sh)
}

destroy_serverless
destroy_terraform
