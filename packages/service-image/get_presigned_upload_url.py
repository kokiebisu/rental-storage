import json
import logging
import os
import boto3
from botocore.exceptions import ClientError

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def handler(event, _):
    response = _get_presigned_upload_url(
        event['queryStringParameters']['filename'])
    # return response
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*',
            'Access-Control-Allow-Headers': 'Content-Type',
            'Access-Control-Allow-Credentials': True,
        },
        'body': json.dumps(response)
    }


def _get_presigned_upload_url(filename: str):
    """Generate a presigned URL S3 POST request to upload a file

    :param bucket_name: string
    :param object_name: string
    :param fields: Dictionary of prefilled form fields
    :param conditions: List of conditions to include in the policy
    :param expiration: Time in seconds for the presigned URL to remain valid
    :return: Dictionary with the following keys:
        url: URL to post to
        fields: Dictionary of form fields and values to submit with the POST
    :return: None if error.
    """
    stage = os.environ['STAGE']
    account_id = os.environ['ACCOUNT_ID']
    access_key_id = os.environ['TERRAFORM_USER_ACCESS_KEY_ID']
    secret_access_key = os.environ['TERRAFORM_USER_SECRET_ACCESS_KEY']
    s3_client = boto3.client('s3', aws_access_key_id=access_key_id,
                             aws_secret_access_key=secret_access_key)
    try:
        return s3_client.generate_presigned_post(
            Bucket=f'{stage}-{account_id}-space-profile',
            Key=filename,
            ExpiresIn=3600
        )
    except ClientError:
        raise
