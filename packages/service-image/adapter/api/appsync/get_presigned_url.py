import logging

from adapter.s3_object_storage_adapter import S3ObjectStorageAdapter
from app.service import ImageService

def handler(event, _):
    logging.info("EVENT: ", event)
    adapter = S3ObjectStorageAdapter()
    service = ImageService(adapter)
    response = service.create_presigned_post('rental-storage-listing-dev-profile', 'random')
    return response
