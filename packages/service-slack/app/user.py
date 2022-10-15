from typing import Any

from domain.error import EventNameNotFoundException
from domain.user import User
from adapter.bot import SlackBotAdapter


class SlackUserMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def send_user_account_event_message(self, event_name: str, message: Any):
        user_account = User(first_name=message['firstName'], 
                    last_name=message['lastName'], 
                    email_address=message['emailAddress']
                )
        if event_name == "created":
            self._alert_user_account_created(user_account)
        else:
            raise EventNameNotFoundException

    def _alert_user_account_created(self, entity: User):
        channel_name = 'C0464PCNZH8'
        message = f'User Signed up :tada:\n>Name: {entity.first_name} {entity.last_name}\n>Email Address: {entity.email_address}'
        self.bot.send_chat_message(channel_name=channel_name, message=message)
