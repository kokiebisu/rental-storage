import json
import logging

from app.adapter.service import SessionService

logger = logging.getLogger(__name__)

def handler(event, context):
    service = SessionService()
    query_params = event['queryStringParameters']
    token = query_params['authorizationToken']
    service.get_login(token)

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
