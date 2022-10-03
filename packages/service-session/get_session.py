import json


def handler(event, context):
    body = {
        "message": "Go Serverless v3.0! Your function executed successfully GET!",
        "input": event,
    }

    response = {"statusCode": 200, "body": json.dumps(body)}

    return response
