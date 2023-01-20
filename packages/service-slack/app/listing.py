from typing import Any

from domain.listing import Listing
from domain.error import EventNameNotFoundException
from adapter.bot import SlackBotAdapter


class SlackListingMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def send_listing_event_message(self, event_name: str, message: Any):
        listing = Listing(
                    lender_id=message['lenderId'],
                    street_address=message['streetAddress']
                )
        if event_name == "created":
            self._alert(entity=listing)
        else:
            raise EventNameNotFoundException

    def _alert(self, entity: Listing):
        channel_name = 'C0464PCNZH8'
        message = SlackListingMessageSenderService._generate_listing_created_message(lender_id=entity["lender_id"], street_address=entity["street_address"])
        self.bot.send_chat_message(channel_name=channel_name, message=message)
    
    def _generate_listing_created_message(self, lender_id: str, street_address: str):
        return (
            f'Listing was posted! :tada:\n'
            f'Lender Id: {lender_id}\n'
            f'Street Address: {street_address}\n'
        )
