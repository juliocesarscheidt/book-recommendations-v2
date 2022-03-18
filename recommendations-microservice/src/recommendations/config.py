import os


def get_amqp_config():
    AMQP_CONN_STRING = os.environ.get("AMQP_CONN_STRING")
    AMQP_QUEUE = os.environ.get("AMQP_QUEUE", "recommendations_queue")
    return {
        "conn_string": AMQP_CONN_STRING,
        "queue": AMQP_QUEUE,
    }


def get_redis_config():
    REDIS_CONN_STRING = os.environ.get("REDIS_CONN_STRING")
    return {
        "conn_string": REDIS_CONN_STRING,
    }


def get_grpc_config():
    GRPC_CONN_STRING = os.environ.get("GRPC_CONN_STRING")
    return {
        "conn_string": GRPC_CONN_STRING,
    }
