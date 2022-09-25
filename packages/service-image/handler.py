
def hello(event, context):
    print("EVENT: ", event)
    print("CONTEXT: ", context)
    return {
        'url': 'some random url'
    }