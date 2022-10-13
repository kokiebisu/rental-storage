import boto3
from slack_sdk.web.client import WebClient

class SlackBotAdapter:
    client: WebClient

    def __init__(self):
        ssm = boto3.client('ssm')
        parameter = ssm.get_parameter(Name='/dev/service-slack/bot-api-key', WithDecryption=True)
        bot_api_key = parameter['Parameter']['Value']
        self.client = WebClient(token=bot_api_key)

    def send_message(self, channel_name: str, message: str):
        self.client.chat_postMessage(
            channel=channel_name,
            text=message
        )