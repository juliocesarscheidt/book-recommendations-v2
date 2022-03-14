from redis import Redis


def get_redis_client(redis_host, redis_port):
    return Redis(
        host=redis_host, port=redis_port, db=0, charset="utf-8", decode_responses=True
    )
