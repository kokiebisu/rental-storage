import boto3

from port.ObjectStorageAdapter import ObjectStorageAdapter


class S3ObjectStorageAdapter(ObjectStorageAdapter):
    _client: boto3.client('s3')

    def __init__(self):
        self._client = boto3.client('s3')