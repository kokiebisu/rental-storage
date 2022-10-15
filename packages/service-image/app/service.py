import logging

from botocore.exceptions import ClientError
from port.ObjectStorageAdapter import ObjectStorageAdapter

class ImageService:
    def __init__(self, object_storage_adapter: ObjectStorageAdapter) -> None:
        self.object_storage_adapter = object_storage_adapter
        
    def create_presigned_post(self, bucket_name, object_name,
                          fields=None, conditions=None, expiration=3600):
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
        
        try:
            response = self.object_storage_adapter.generate_presigned_post(bucket_name,
                                                        object_name,
                                                        Fields=fields,
                                                        Conditions=conditions,
                                                        ExpiresIn=expiration)
        except ClientError as e:
            logging.error(e)
            return None

        # The response contains the presigned URL and required fields
        return response
