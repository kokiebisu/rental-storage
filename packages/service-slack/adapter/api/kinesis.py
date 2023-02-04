from slack_sdk.errors import SlackApiError

from domain.error import (EntityTypeNotFoundException,
                          EventNameNotFoundException)
from adapter.stream_event_message_handler import KinesisEventStreamHandler
from app.user import SlackUserMessageSenderService
from app.listing import SlackListingMessageSenderService
from app.booking import SlackBookingMessageSenderService
from adapter.bot import SlackBotAdapter


def handler(event, context):
    messages = KinesisEventStreamHandler.parse_event(event)
    bot_adapter = SlackBotAdapter()

    try:
        for message in messages:
            if 'data' in message:
                if message['eventEntity'] == 'User':
                    service = SlackUserMessageSenderService(bot_adapter)
                    service\
                        .send_user_account_event_message(
                            event_name=message['eventName'],
                            message=message['data'])
                if message['eventEntity'] == 'Listing':
                    service = SlackListingMessageSenderService(bot_adapter)
                    service.send_listing_event_message(message['eventName'],
                                                       message['data'])
                if message['eventEntity'] == 'Booking':
                    service = SlackBookingMessageSenderService(bot_adapter)
                    service.send_booking_event_message(message['eventName'],
                                                       message['data'])
                else:
                    raise EntityTypeNotFoundException('Unavailable entityType')
    except EventNameNotFoundException:
        return {
            "statusCode": 500,
            "body": "Unavailable event name"
        }
    except EntityTypeNotFoundException:
        return {
            "statusCode": 500,
            "body": "Unavailable entity type"
        }
    except SlackApiError:
        return {
            "statusCode": 500,
            "body": "Something went wrong with the Slack API"
        }
    else:
        return {
            "statusCode": 200,
        }
