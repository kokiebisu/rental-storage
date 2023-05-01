import json
import boto3


import os


def handler(event, context):

    print("EVENT: ", event)
    print("***")
    print("CONTEXT: ", context)
    domain = event['requestContext']['domainName']
    endpoint_url = f'https://{domain}/local'

    connectionId = event['requestContext']['connectionId']

    apigateway = boto3.client('apigatewaymanagementapi',
                              endpoint_url=endpoint_url)
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table('service-chat-connection-ids')

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
        print("ELSE!")
        apigateway.post_to_connection(
            ConnectionId=connectionId,
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
