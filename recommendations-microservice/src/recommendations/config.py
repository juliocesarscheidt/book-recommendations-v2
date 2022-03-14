import os


def get_amqp_config():
    AMQP_CONN_STRING = os.environ.get("AMQP_CONN_STRING")
    AMQP_QUEUE = os.environ.get("AMQP_QUEUE", "recommendations_queue")
    return {
        "conn_string": AMQP_CONN_STRING,
        "queue": AMQP_QUEUE,
    }


def get_redis_config():
    REDIS_HOST = os.environ.get("REDIS_HOST")
    REDIS_PORT = os.environ.get("REDIS_PORT")
    return {
        "host": REDIS_HOST,
        "port": REDIS_PORT,
    }


def get_api_gateway_config():
    API_GATEWAY_CONN_STRING = os.environ.get("API_GATEWAY_CONN_STRING")
    return {
        "uri": API_GATEWAY_CONN_STRING,
    }
