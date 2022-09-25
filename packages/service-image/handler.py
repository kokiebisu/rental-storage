import service

def get_presigned_url(event, context):
    print("EVENT: ", event)
    print("Context: ", context)
    response = service.create_presigned_post('rental-storage-listing-dev-profile', 'random')
    return response
