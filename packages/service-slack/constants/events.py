from enum import Enum


class UserEvents(Enum):
    USER_CREATED = 'user_created'
    USER_DELETED = 'user_deleted'
    USER_BOOKED = 'user_booked'


class ListingEvents(Enum):
    LISTING_CREATED = 'listing_created'


class BookingEvents(Enum):
    BOOKING_CREATED = 'booking_created'
