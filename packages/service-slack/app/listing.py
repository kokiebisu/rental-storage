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
            self._alert_listing_created(listing)
        else:
            raise EventNameNotFoundException

    def _alert_listing_created(self, entity: Listing):
        channel_name = 'C0464PCNZH8'
        message = f'Listing was posted! :tada:\n>Lender Id: {entity.lender_id}\n>Street Address: {entity.street_address}'
        self.bot.send_chat_message(channel_name=channel_name, message=message)
