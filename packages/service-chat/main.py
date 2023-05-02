from datetime import datetime

import json
import boto3
import logging

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def handler(event, context):
    logger.info("Entered lambda handler")
    print("EVENT: ", event)
    print("***")
    print("CONTEXT: ", context)
    domain = event['requestContext']['domainName']
    endpoint_url = f'https://{domain}/local'

    connection_id = event['requestContext']['connectionId']

    apigateway = boto3.client('apigatewaymanagementapi',
                              endpoint_url=endpoint_url)
    dynamodb = boto3.resource('dynamodb')
    tableName = '{}-chat-connections'.format('local')
    table = dynamodb.Table(tableName)

    request_context = event['requestContext']
    if request_context['routeKey'] == '$connect':
        user_id = event['queryStringParameters']['userID']
        destination_user_id = event['queryStringParameters']['destinationUserID']
        item = {
            'UserId': user_id,
            'DestinationUserId': destination_user_id,
            'ConnectionId': connection_id,
            'CreatedAt': int(datetime.utcnow().timestamp())
        }
        logger.info('Added connection to table')
        table.put_item(Item=item)
        return {
            'statusCode': 200,
        }

    if request_context['routeKey'] == '$disconnect':
        table.delete_item(Key={
            'ConnectionId': connection_id
        })
        logger.info('Removed connection from table')
        return {
            'statusCode': 200,
        }

    else:
        print("ELSE!")
        apigateway.post_to_connection(
            ConnectionId=connection_id,
            Data=json.dumps({
                'message': 'Hello from the other side',
            }).encode('utf-8')
        )

        return {
            'statusCode': 200,
        }

    # table.put_item(Item={
    #     'connectionId': connection_id,
    #     'message': message,
    #     'sender': sender
    # })
