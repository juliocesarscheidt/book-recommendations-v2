from threading import Thread

from application.service.service import Service
from infra.adapter.amqp_adapter import get_connection, get_channel
from infra.adapter.publisher import Publisher
from infra.adapter.cache import Cache
from infra.adapter.grpc_client import GrpcClient


class Controller(object):
    publisher: Publisher
    cache: Cache
    service: Service
    grpc_client: None

    def __init__(self, redis_client, grpc_conn_string) -> None:
        self.publisher = Publisher()
        self.cache = Cache(redis_client)
        self.grpc_client = GrpcClient(grpc_conn_string)
        self.service = Service(self.publisher, self.cache, self.grpc_client)

    def calculate(self, __message, __properties, __amqp_parameters):
        __connection = get_connection(__amqp_parameters)
        __channel = get_channel(__connection)
        # get rates
        thread = Thread(
            target=self.service.calculate_rates,
            args=[__message, __channel, __properties],
        )
        thread.start()
        return True

    def get(self, __message, __properties, __amqp_parameters):
        __connection = get_connection(__amqp_parameters)
        __channel = get_channel(__connection)
        # get recommendation
        thread = Thread(
            target=self.service.get_recommendation,
            args=[__message, __channel, __properties],
        )
        thread.start()
        return True
