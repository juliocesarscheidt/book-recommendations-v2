from redis import Redis

from application.adapter.abstract_redis_adapter import AbstractRedisAdapter


class RedisAdapter(AbstractRedisAdapter):
    redis_conn_string: None
    redis_client: None

    def __init__(self, redis_conn_string) -> None:
        self.redis_conn_string = redis_conn_string
        self.redis_client = self.get_redis_client()

    def get_redis_client(self):
        return Redis.from_url(
            url=f"redis://{self.redis_conn_string}/0",
            charset="utf-8",
            decode_responses=True,
        )

    def get_cache(self, key):
        cached = self.redis_client.get(key)
        return cached

    def set_cache(self, key, data, hours=24) -> None:
        exp = 60 * 60 * hours  # default 24 hours
        self.redis_client.setex(key, exp, data)
