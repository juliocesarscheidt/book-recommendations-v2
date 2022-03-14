from threading import Thread

from infra.adapter.amqp_adapter import get_connection, get_channel
from infra.adapter.publisher import Publisher
from infra.adapter.cache import Cache
from application.service.service import Service


class Controller(object):
    publisher: Publisher
    cache: Cache
    service: Service
    api_gateway_uri: None

    def __init__(self, redis_client, api_gateway_uri) -> None:
        self.publisher = Publisher()
        self.cache = Cache(redis_client)
        self.api_gateway_uri = api_gateway_uri
        self.service = Service(self.publisher, self.cache, self.api_gateway_uri)

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
