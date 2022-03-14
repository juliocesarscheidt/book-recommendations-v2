class AbstractCache(object):
    redis_client: None

    def __init__(self, redis_client) -> None:
        pass

    def get_cache(self, key):
        pass

    def set_cache(self, key, data, hours=24) -> None:
        pass
