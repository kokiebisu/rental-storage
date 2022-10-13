import logging
from slack_sdk.errors import SlackApiError

from domain.error import EntityTypeNotFoundException, EventNameNotFoundException
from adapter.event_message_handler import SQSEventMessageHandler
from app.user import SlackUserMessageSenderService
from adapter.bot import SlackBotAdapter

def handler(event, context):
    print("EVENT: ", event)
    print("CONTEXT: ", context)
    logging.info("EVENT: ", event)

    messages = SQSEventMessageHandler.parse_event(event)
    print("ENTERED1")
    bot_adapter = SlackBotAdapter()
    service = SlackUserMessageSenderService(bot_adapter)
    try:
        print("ENTERED2")
        for message in messages:
            if 'data' in message:
                print("MESSAGE: ", message)
                if message['entityType'] == 'user_account':
                    service.send_user_account_event_message(message['eventName'], message['data'])
                else:
                    raise EntityTypeNotFoundException('Unavailable entityType')
    except EventNameNotFoundException as e:
        return {
            "statusCode": 500,
            "body": "Unavailable event name"
        }
    except EntityTypeNotFoundException as e:
        return {
            "statusCode": 500,
            "body": "Unavailable entity type"
        }
    except SlackApiError as e:
        return {
            "statusCode": 500,
            "body": "Something went wrong with the Slack API"
        }
    else:
        return {
            "statusCode": 200,
        }

  
