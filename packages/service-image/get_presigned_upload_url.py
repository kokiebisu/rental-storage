import json
import logging
import os
import boto3
from botocore.exceptions import ClientError

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def handler(event, context):
    logging.info("EVENT: " + str(event))
    print("EVENT: ", event)
    # THIS NEED TO BE CONVERTED TO REST
    response = _get_presigned_upload_url(json.loads(event['body'])['filename'])
    # return response
    return {
        'statusCode': 200,
        'headers': {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
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
    print("ENTERED1: ", filename)
    # Generate a presigned S3 POST URL
    stage = os.environ['STAGE']
    account_id = os.environ['ACCOUNT_ID']
    s3_client = boto3.client('s3', region_name="us-east-1",
                             config=boto3.session.Config(signature_version='s3v4'))
    try:
        put_url = s3_client.generate_presigned_url(
            'put_object',
            Params={
                'Bucket': f'{stage}-{account_id}-listing-profile',
                'Key': filename,
                'ACL': 'public-read'
            })
        print("Got presigned POST URL: %s", put_url)
    except ClientError:
        print(
            "Couldn't get a presigned POST URL for bucket '%s' and object '%s'",
            "dev-app-listing-profile", put_url)
        raise
    return {
        'url': put_url,
        'filename': filename
    }
