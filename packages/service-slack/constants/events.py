from enum import Enum


class UserEvents(Enum):
    USER_CREATED = 'user_created'
    USER_DELETED = 'user_deleted'


class SpaceEvents(Enum):
    SPACE_CREATED = 'space_created'
    SPACE_DELETED = 'space_deleted'


class BookingEvents(Enum):
    BOOKING_CREATED = 'booking_created'
