class AbstractRedisAdapter(object):
    redis_client: None

    def __init__(self, redis_conn_string) -> None:
        pass

    def get_redis_client(self, redis_conn_string):
        pass

    def get_cache(self, key):
        pass

    def set_cache(self, key, data, hours=24) -> None:
        pass
