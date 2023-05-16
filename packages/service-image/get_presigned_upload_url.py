import json
import logging
import os
import boto3
from botocore.exceptions import ClientError

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def handler(event, _):
    print("EVENT: ", event)
    # THIS NEED TO BE CONVERTED TO REST
    response = _get_presigned_upload_url(json.loads(event['body'])['filename'])
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
    # Generate a presigned S3 POST URL
    stage = os.environ['STAGE']
    account_id = os.environ['ACCOUNT_ID']
    s3_client = boto3.client('s3', region_name="us-east-1",
                             config=boto3.session.Config(signature_version='s3v4'))
    key = f'{stage}/{filename}'
    try:
        put_url = s3_client.generate_presigned_url(
            'put_object',
            Params={
                'Bucket': f'{stage}-{account_id}-space-profile',
                'Key': key,
            },
            ExpiresIn=3600)
        print("Got presigned POST URL: %s", put_url)
    except ClientError:
        raise
    return {
        'presignedUrl': put_url,
    }
