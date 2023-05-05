from datetime import datetime

import json
import boto3
import logging
import uuid

logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def handler(event, context):
    print("EVENT: ", event)
    domain = event['requestContext']['domainName']
    endpoint_url = f'https://{domain}/local'

    # set up apis
    dynamodb = boto3.client('dynamodb')
    apigateway = boto3.client('apigatewaymanagementapi',
                              endpoint_url=endpoint_url)

    connections_table_name = '{}-chat-connections'.format('local')
    messages_table_name = '{}-chat-messages'.format('local')

    connection_id = event['requestContext']['connectionId']

    request_context = event['requestContext']

    if request_context['routeKey'] == '$connect':
        sender_id = event['queryStringParameters']['senderId']
        recipient_id = event['queryStringParameters']['recipientId']
        register_connection(client=dynamodb, connections_table_name=connections_table_name, sender_id=sender_id,
                            recipient_id=recipient_id, connection_id=connection_id)

        # display message history between the sender and recipient
        messages = retrieve_message_history(
            client=dynamodb, messages_table_name=messages_table_name, sender_id=sender_id, recipient_id=recipient_id)

        if len(messages):
            apigateway.post_to_connection(
                ConnectionId=connection_id,
                Data=json.dumps({
                    'messages': messages
                }).encode('utf-8')
            )

    elif request_context['routeKey'] == '$disconnect':
        unregister_connection(
            client=dynamodb, connections_table_name=connections_table_name, connection_id=connection_id)

    else:
        message = json.loads(event['body'])['message']
        send_message(dynamodb_client=dynamodb, apigateway_client=apigateway,
                     connections_table_name=connections_table_name,
                     messages_table_name=messages_table_name,
                     sender_connection_id=connection_id,
                     message=message)

    return {
        'statusCode': 200,
    }


def retrieve_message_history(client, messages_table_name, sender_id, recipient_id):
    logger.info('retrieve_message_history')
    response = client.query(
        TableName=messages_table_name,
        IndexName='ChatMessagesSenderIdRecipientIdIndex',
        KeyConditionExpression='RecipientId = :recipient_id AND SenderId = :sender_id',
        ExpressionAttributeValues={
            ':sender_id': {'S': sender_id},
            ':recipient_id': {'S': recipient_id}
        }
    )
    print("RESPONSE: ", response)
    return response['Items']


def retrieve_recipient_connection_id(client, connections_table_name, sender_id, recipient_id):
    # logger.info('retrieve_recipient_connection_id', recipient_id, sender_id)
    response = client.query(
        TableName=connections_table_name,
        IndexName='ChatConnectionsSenderIdRecipientIdIndex',
        KeyConditionExpression='RecipientId = :recipient_id AND SenderId = :sender_id',
        ExpressionAttributeValues={
            ':sender_id': {'S': recipient_id},
            ':recipient_id': {'S': sender_id}
        }
    )
    item = response['Items'][0]
    return item['ConnectionId']['S']


def register_connection(client, connections_table_name, sender_id, recipient_id, connection_id):
    logger.info('register_connection')
    item = {
        'SenderId': {'S': sender_id},
        'RecipientId': {'S': recipient_id},
        'ConnectionId': {'S': connection_id},
        'CreatedAt': {'S': str(datetime.utcnow().timestamp())}
    }
    client.put_item(TableName=connections_table_name, Item=item)


def unregister_connection(client, connections_table_name, connection_id):
    """
    Unregister connection id from db table
    """
    logger.info('unregister_connection')
    client.delete_item(TableName=connections_table_name, Key={
        'ConnectionId': {'S': connection_id}
    })


def retrieve_by_sender_connection_id(client, connections_table_name, connection_id):
    """
    Retrieve the sender and recipient ids by the sender connection id
    """
    logger.info('retrieve_by_sender_connection_id')
    response = client.get_item(
        TableName=connections_table_name,
        Key={
            'ConnectionId': {'S': connection_id}
        }
    )
    item = response['Item']
    return {
        'senderId': item['SenderId']['S'],
        'recipientId': item['RecipientId']['S']
    }


def send_message(dynamodb_client, apigateway_client, connections_table_name, messages_table_name, sender_connection_id, message):
    # retrieve the sender id and recipient id
    connection_data = retrieve_by_sender_connection_id(
        client=dynamodb_client, connections_table_name=connections_table_name, connection_id=sender_connection_id)

    # save the message to db
    dynamodb_client.put_item(TableName=messages_table_name, Item={
        'UId': {'S': uuid.uuid1()},
        'SenderId': {'S': connection_data['senderId']},
        'RecipientId': {'S': connection_data['recipientId']},
        'Message': {'S': message},
        'CreatedAt': {'S': str(datetime.utcnow().timestamp())}
    })

    # retrieve the recipient connection id for the recipient
    recipient_connection_id = retrieve_recipient_connection_id(
        client=dynamodb_client, connections_table_name=connections_table_name,
        sender_id=connection_data['senderId'], recipient_id=connection_data['recipientId'])
    # send the message to the recipient
    apigateway_client.post_to_connection(
        ConnectionId=recipient_connection_id,
        Data=json.dumps({
            'message': message
        }).encode('utf-8')
    )
