from redis import Redis


def get_redis_client(redis_conn_string):
    return Redis.from_url(
        url=f"redis://{redis_conn_string}/0", charset="utf-8", decode_responses=True
    )
