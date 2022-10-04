import json
from app.adapter.service import SessionService


def handler(event, context):
    service = SessionService()
    body = json.loads(event['body'])
    print(body)
    # logger.info("QUERY PARAMS: ", query_params)
    token = body["authorizationToken"]
    user = body['user']
    service.set_login(token, user)
    return {
        "statusCode": 200,
        "message": "worked"
    }
