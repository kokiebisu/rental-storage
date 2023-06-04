#!/bin/bash

cleanup() {
  # Terminate SSH processes
  if [[ -n "$user_db_pid" ]]; then
    kill "$user_db_pid"
  fi
  if [[ -n "$space_db_pid" ]]; then
    kill "$space_db_pid"
  fi

  exit 0
}

# Set cleanup function to execute on SIGINT
trap cleanup SIGINT

# Read JSON files and assign variables
SSH_KEY_NAME=id_rsa_rental_storage_ec2
USER_DB_ENDPOINT=$(jq -r '.user' terraform/json/rds.json)
SPACE_DB_ENDPOINT=$(jq -r '.space' terraform/json/rds.json)
BASTION_IP=$(jq -r '.endpoint' terraform/json/ec2_bastion.json)


# Start SSH sessions
ssh -i ~/.ssh/${SSH_KEY_NAME} -N -L localhost:5432:${USER_DB_ENDPOINT} ubuntu@${BASTION_IP} &
user_db_pid=$!

ssh -i ~/.ssh/${SSH_KEY_NAME} -N -L localhost:5433:${SPACE_DB_ENDPOINT} ubuntu@${BASTION_IP} &
space_db_pid=$!

# Wait for script termination
wait

# Cleanup after script termination
cleanup
