import logging
from app.adapter.store import CacheStore

logger = logging.getLogger(__name__)

class SessionService:
    _store: CacheStore
    def __init__(self):
        self._store = CacheStore()

    def get_login(self, token) -> str:
        # logger.info("SessionService|get_login()", token)
        print("get_login() ", token)
        return self._store.get_token('login:', token)

    def set_login(self, token: str, user: any) -> None:
        # logger.info("SessionService|set_login()", token)
        print("set_login() ", token, user)
        self._store.set_token('login:', token, user)
