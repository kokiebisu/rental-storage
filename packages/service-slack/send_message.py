import logging
import boto3
from slack_sdk.errors import SlackApiError
from slack_sdk.web.client import WebClient

def handler(event, _):
    logging.info("EVENT: ", event)
    ssm = boto3.client('ssm')
    parameter = ssm.get_parameter(Name='/dev/service-slack/bot-api-key', WithDecryption=True)
    bot_api_key = parameter['Parameter']['Value']

    client = WebClient(token=bot_api_key)

    try:
        client.chat_postMessage(
            channel="C0464PCNZH8",
            text="Hello from your app! :tada:"
        )
    except SlackApiError as e:
        return {
            "statusCode": 500,
            "body": e
        }

    return {
        "statusCode": 200,
        "body": "hello"
    }
