import pika

from config import get_amqp_config, get_redis_config, get_grpc_config
from infra.handler.event_handler import EventHandler

if __name__ in "__main__":
    amqp_config = get_amqp_config()
    redis_config = get_redis_config()
    grpc_config = get_grpc_config()

    event_handler = EventHandler(
        pika.URLParameters(amqp_config["conn_string"]),
        amqp_config["queue"],
        redis_config["conn_string"],
        grpc_config["conn_string"],
    )
    event_handler.start()
