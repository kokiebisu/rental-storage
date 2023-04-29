import json
import boto3

# dynamodb = boto3.resource('dynamodb')
# table = dynamodb.Table('ChatMessages')

import os


def handler(event, context):

    print("EVENT: ", event)
    print("HEADERS: ", event['headers'])
    print("***")
    print("CONTEXT: ", context)
    domain = event['headers']['Host']
    endpoint_url = f'https://{domain}/local'
    apigateway = boto3.client('apigatewaymanagementapi',
                              endpoint_url=endpoint_url)

    print("ENDPOINT_URL: ", endpoint_url)

    request_context = event['requestContext']
    if request_context['routeKey'] == '$connect':
        print('CONNECTED')
        return {
            'statusCode': 200,
        }

    if request_context['routeKey'] == '$disconnect':
        print('DISCONNECTED')
        return {
            'statusCode': 200,
        }

    else:
        print("ENTERED")
        apigateway.post_to_connection(
            ConnectionId=event['requestContext']['connectionId'],
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
