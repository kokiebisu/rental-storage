import logging
from slack_sdk.errors import SlackApiError

from app.user import SlackUserMessageSenderService
from adapter.bot import SlackBotAdapter

def handler(event, _):
    logging.info("EVENT: ", event)
    
    bot_adapter = SlackBotAdapter()
    service = SlackUserMessageSenderService(bot_adapter)

    try:
        # check the name of topic and trigger if it matches
        if True:
            service.alert_user_sign_up()
    except SlackApiError as e:
        return {
            "statusCode": 500,
            "body": e.response
        }

    return {
        "statusCode": 200,
    }
