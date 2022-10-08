import json
import os
from app.adapter.service import SessionService
from redis import Redis


def handler(event, context):
    # service = SessionService()
    host = os.environ['REDIS_HOST']
    print(host)
    port = os.environ['REDIS_PORT']
    print(port)
    redis = Redis(host=host, port=port, decode_responses=True, ssl=True, username='myuser', password='MyPassword0123456789')
    print('ready to ping')
    if redis.ping():
        print('Connected using ping')
    # body = json.loads(event['body'])
    # print(body)
    # logger.info("QUERY PARAMS: ", query_params)
    # token = body["authorizationToken"]
    # user = body['user']
    redis.set('hello', 'world')
    # service.set_login(token, user)
    return {
        "statusCode": 200,
        "message": "worked"
    }
