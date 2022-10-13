import logging
from slack_sdk.errors import SlackApiError

from domain.error import SourceServiceNotFoundException
from adapter.event_message_handler import SQSEventMessageHandler
from app.user import SlackUserMessageSenderService
from adapter.bot import SlackBotAdapter

def handler(event, _):
    logging.info("EVENT: ", event)

    messages = SQSEventMessageHandler.parse(event)
    bot_adapter = SlackBotAdapter()
    service = SlackUserMessageSenderService(bot_adapter)
    try:
        for message in messages:
            if 'sourceService' in message:
                if message['sourceService'] == 'user':
                    service.alert_user_sign_up(message)
                else:
                    raise SourceServiceNotFoundException('Unavailable sourceService')

    except SlackApiError as e:
        return {
            "statusCode": 500,
            "body": e.response
        }
    else:
        return {
            "statusCode": 200,
        }

  
