from application.adapter.abstract_cache import AbstractCache


class Cache(AbstractCache):
    redis_client: None

    def __init__(self, redis_client) -> None:
        self.redis_client = redis_client

    def get_cache(self, key):
        cached = self.redis_client.get(key)
        return cached

    def set_cache(self, key, data, hours=24) -> None:
        exp = 60 * 60 * hours  # default 24 hours
        self.redis_client.setex(key, exp, data)
