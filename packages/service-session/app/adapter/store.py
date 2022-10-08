import os
from redis import StrictRedis
import logging

logging.basicConfig(level=logging.INFO)

class CacheStore:
    _client: StrictRedis

    def __init__(self):
        host = os.getenv("REDIS_HOST")
        port = os.getenv("REDIS_PORT")
        self._client = StrictRedis(host=host, port=port, db=0, ssl=True)
        print(f"self._client: {self._client}")

    def get_token(self, category, token):
        print("get_token: ", category, token)
        result = self._client.hget(category, token)
        print(f"RESULT: {result}")
        return result

    def set_token(self, category, token, user):
        print("set_token: ", category, token, user)
        self._client.hset(category, token, user)
        print("SUCCESS")