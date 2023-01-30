from typing import Any

from constants.events import UserEvents
from domain.error import EventNameNotFoundException
from domain.user import User
from adapter.bot import SlackBotAdapter


class SlackUserMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def send_user_account_event_message(self, event_name: str, message: Any):
        channel_name = SlackUserMessageSenderService.get_channel_name()
        user_account = User(first_name=message['firstName'],
                            last_name=message['lastName'],
                            email_address=message['emailAddress'])
        if event_name in UserEvents:
            if event_name == UserEvents.USER_CREATED:
                message = SlackUserMessageSenderService\
                    .generate_user_account_created_message(
                        event_name=event_name,
                        first_name=user_account.first_name,
                        last_name=user_account.last_name,
                        email_address=user_account.email_address)
            elif event_name == UserEvents.USER_DELETED:
                message = SlackUserMessageSenderService\
                    .generate_user_account_deleted_message(
                        event_name=event_name,
                        first_name=user_account.first_name,
                        last_name=user_account.last_name,
                        email_address=user_account.email_address)
            return self.bot.send_chat_message(channel_name=channel_name,
                                              message=message)
        else:
            raise EventNameNotFoundException

    @classmethod
    def get_channel_name():
        return 'C0464PCNZH8'

    @classmethod
    def generate_user_account_created_message(cls, first_name,
                                              last_name, email_address):
        return (
            f'User Signed up! :tada:\n'
            f'Name: {first_name} {last_name}\n'
            f'Email Address: {email_address}'
        )

    @classmethod
    def generate_user_account_deleted_message(cls, first_name,
                                              last_name, email_address):
        return (
            f'User has deleted the account! :tada:\n'
            f'Name: {first_name}{last_name}\n'
            f'Email Address: {email_address}\n'
        )
