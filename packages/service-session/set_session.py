import logging
import service

def handler(event, _):
    logging.info("EVENT: ", event)
    response = service.create_presigned_post('rental-storage-listing-dev-profile', 'random')
    return response
