from typing import Any

from constants.events import BookingEvents
from domain.error import EventNameNotFoundException
from domain.booking import Booking
from adapter.bot import SlackBotAdapter


class SlackBookingMessageSenderService:
    bot: SlackBotAdapter

    def __init__(self, bot_adapter: SlackBotAdapter):
        self.bot = bot_adapter

    def send_booing_event_message(self, event_name: str, message: Any):
        channel_name = SlackBookingMessageSenderService.get_channel_name()
        booking = Booking(
                    first_name=message['firstName'],
                    last_name=message['lastName'],
                    email_address=message['emailAddress'],
                    user_id=message['userId'],
                )
        if event_name in BookingEvents:
            message = SlackBookingMessageSenderService\
                .generate_user_booking_created_message(
                    event_name=event_name,
                    first_name=booking.first_name,
                    last_name=booking.last_name,
                    email_address=booking.email_address,
                    user_id=booking.user_id, )
            self.bot.send_chat_message(channel_name=channel_name,
                                       message=message)
        else:
            raise EventNameNotFoundException

    @classmethod
    def get_channel_name():
        return 'C0464PCNZH8'

    @classmethod
    def generate_booking_created_message(clr, first_name, last_name,
                                         email_address, user_id):
        message = (
            f'User booking created! :tada:\n'
            f'Name: {first_name} {last_name}\n'
            f'Email Address: {email_address}\n'
            f'User Id: {user_id}\n'
        )
        return message
