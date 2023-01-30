from typing import Any

from constants.events import ListingEvents
from domain.listing import Listing
from domain.error import EventNameNotFoundException
from adapter.bot import SlackBotAdapter


class SlackListingMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def send_listing_event_message(self, event_name: str, message: Any):
        channel_name = SlackListingMessageSenderService.get_channel_name()
        listing = Listing(
                    lender_id=message['lenderId'],
                    street_address=message['streetAddress']
                )
        if event_name in ListingEvents:
            if event_name == ListingEvents.LISTING_CREATED:
                message = SlackListingMessageSenderService\
                    .generate_listing_created_message(
                        lender_id=listing.lender_id,
                        street_address=listing.street_address)
            self.bot.send_chat_message(channel_name=channel_name,
                                       message=message)
        else:
            raise EventNameNotFoundException

    @classmethod
    def get_channel_name(cls):
        return 'C0464PCNZH8'

    @classmethod
    def generate_listing_created_message(cls, lender_id: str,
                                         street_address: str):
        return (
            f'Listing was posted! :tada:\n'
            f'Lender Id: {lender_id}\n'
            f'Street Address: {street_address}\n'
        )
