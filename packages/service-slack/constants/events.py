from enum import Enum


class UserEvents(Enum):
    USER_CREATED = 'user_created'
    USER_DELETED = 'user_deleted'


class ListingEvents(Enum):
    LISTING_CREATED = 'listing_created'
