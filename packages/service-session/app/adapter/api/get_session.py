import json
import logging
import os
import sys
sys.path.insert(0, 'vendor')
from redis import Redis

logger = logging.getLogger(__name__)

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

    # # # logger.info("QUERY PARAMS: ", query_params)
 
    # # # logger.info("TOKEN: ", token)
    # user = service.get_login(token)
    # # print("ENTERED4", user)
    # if not user:
    #     return {
    #         "statusCode": 404, 
    #         "body": "User not found"
    #     }
    # # print("ENTERED5")
    # # return {
    # #     "statusCode": 200, 
    # #     "body": json.dumps(user)
    # # }
    return {
        "statusCode": 200,
        "body": json.dumps(token)
    }
