import pika
import logging

from config import get_amqp_config, get_redis_config, get_api_gateway_config
from infra.handler.handler import Handler

logging.basicConfig(format="%(asctime)s - %(message)s", level=logging.INFO)

if __name__ in "__main__":
    amqp_config = get_amqp_config()
    redis_config = get_redis_config()
    api_gateway_config = get_api_gateway_config()

    handler = Handler(
        pika.URLParameters(amqp_config["conn_string"]),
        amqp_config["queue"],
        redis_config["host"],
        redis_config["port"],
        api_gateway_config["uri"],
    )
    handler.handle()
