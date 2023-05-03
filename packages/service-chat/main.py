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
    dynamodb = boto3.client('dynamodb')
    connections_table_name = '{}-chat-connections'.format('local')

    request_context = event['requestContext']
    if request_context['routeKey'] == '$connect':
        user_id = event['queryStringParameters']['userID']
        destination_user_id = event['queryStringParameters']['destinationUserID']
        item = {
            'UserId': {'S': user_id},
            'DestinationUserId': {'S': destination_user_id},
            'ConnectionId': {'S': connection_id},
            'CreatedAt': {'S': str(datetime.utcnow().timestamp())}
        }
        logger.info('Added connection to table')
        dynamodb.put_item(TableName=connections_table_name, Item=item)
        return {
            'statusCode': 200,
        }

    if request_context['routeKey'] == '$disconnect':
        dynamodb.delete_item(TableName=connections_table_name, Key={
            'ConnectionId': connection_id
        })
        logger.info('Removed connection from table')
        return {
            'statusCode': 200,
        }

    else:
        logger.info('Sending message')
        # retrieve the destination user id by the connection id
        response = dynamodb.get_item(
            TableName=connections_table_name,
            Key={
                'ConnectionId': {'S': connection_id}
            }
        )
        item = response['Item']
        destination_user_id = item['DestinationUserId']['S']
        user_id = item['UserId']['S']

        # retrieve the connection id for the destination id
        response = dynamodb.query(
            TableName=connections_table_name,
            IndexName='ChatUserIdDestinationIdIndex',
            KeyConditionExpression='DestinationUserId = :destination_user_id AND UserId = :user_id',
            ExpressionAttributeValues={
                ':user_id': {'S': destination_user_id},
                ':destination_user_id': {'S': user_id}
            }
        )
        item = response['Items'][0]
        destination_connection_id = item['ConnectionId']['S']
        message = json.loads(event['body'])['message']
        apigateway.post_to_connection(
            ConnectionId=destination_connection_id,
            Data=json.dumps({
                'message': message
            }).encode('utf-8')
        )

        return {
            'statusCode': 200,
        }
