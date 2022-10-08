import os
import sys
sys.path.insert(0, 'src/vendor')
from redis import Redis

def handler(event, context):
    # service = SessionService()
    host = os.environ['REDIS_HOST']
    print(host)
    port = os.environ['REDIS_PORT']
    print(port)
    redis = Redis(host=host, port=port, ssl=True)
    query_params = event['queryStringParameters']
    token = query_params['authorizationToken']
    # service.get_login(token)
    value = redis.get('hello')

    if not token:
        return {
            "statusCode": 400,
            "body": "can't return anything"
        }

    return {
        "statusCode": 200,
        "body": "success"
    }