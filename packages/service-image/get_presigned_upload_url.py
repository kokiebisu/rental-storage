import json
import logging
import os
import boto3
from botocore.exceptions import ClientError

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def handler(event, _):
    logger.info(event)
    response = _get_presigned_upload_url()
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


def _get_presigned_upload_url():
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
    def generate_hash():
        """Generate a hash to be used as filename"""
        import hashlib
        import uuid
        hash_object = hashlib.sha1(uuid.uuid4().bytes)
        hex_dig = hash_object.hexdigest()
        return hex_dig
    hash = generate_hash()
    stage = os.environ['STAGE']
    account_id = os.environ['ACCOUNT_ID']
    access_key_id = os.environ['PRESIGNED_URL_ALLOW_ACCESS_KEY_ID']
    secret_access_key = os.environ['PRESIGNED_URL_ALLOW_SECRET_ACCESS_KEY']
    s3_client = boto3.client('s3', aws_access_key_id=access_key_id,
                             aws_secret_access_key=secret_access_key)
    try:
        presigned_url = s3_client.generate_presigned_post(
            Bucket=f'{stage}-{account_id}-space-profile',
            Key=hash,
            ExpiresIn=3600
        )
        return {
            'presignedUrl': presigned_url,
            'key': hash
        }
    except ClientError:
        raise
