#!/bin/bash

# Set the AWS region
export AWS_DEFAULT_REGION="us-east-1"
export ENV="local"

function remove_ecs_service_images () {
    package=$1
    # Iterate over the repositories

    repository_name="ecs-service-${package}-${ENV}"
    echo "Cleaning up images in repository: $repository_name"

    # Get the list of image IDs in the repository
    image_ids=$(aws ecr list-images --repository-name "$repository_name" --query 'imageIds[].imageDigest' --output text)

    # Iterate over the image IDs and delete each image
    for image_id in $image_ids; do
        aws ecr batch-delete-image --repository-name "$repository_name" --image-ids "imageDigest=$image_id"
        echo "Deleted image: $image_id"
    done

}

remove_ecs_service_images $1