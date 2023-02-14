from typing import Any

from constants.events import SpaceEvents
from domain.space import Space
from domain.error import EventNameNotFoundException
from adapter.bot import SlackBotAdapter


class SlackSpaceMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def send_space_event_message(self, event_name: str, message: Any):
        channel_name = SlackSpaceMessageSenderService.get_channel_name()
        space = Space(
                    lender_id=message['lenderId'],
                    street_address=message['streetAddress']
                )
        if event_name in SpaceEvents:
            message = SlackSpaceMessageSenderService\
                .generate_space_created_message(
                    lender_id=space.lender_id,
                    street_address=space.street_address)
            self.bot.send_chat_message(channel_name=channel_name,
                                       message=message)
        else:
            raise EventNameNotFoundException

    @classmethod
    def get_channel_name(cls):
        return 'C0464PCNZH8'

    @classmethod
    def generate_space_created_message(cls, lender_id: str,
                                       street_address: str):
        return (
            f'Space was posted! :tada:\n'
            f'Lender Id: {lender_id}\n'
            f'Street Address: {street_address}\n'
        )

    @classmethod
    def generate_space_deleted_message(cls, lender_id: str,
                                       street_address: str):
        return (
            f'Space was deleted! :tada:\n'
            f'Lender Id: {lender_id}\n'
            f'Street Address: {street_address}\n'
        )
