import json
import logging

from infra.adapter.redis_adapter import get_redis_client
from infra.adapter.amqp_adapter import get_connection, get_channel
from infra.controller.controller import Controller
from application.commands.commands import Commands


class Handler(object):
    controller: Controller
    commands: Commands
    amqp_parameters: None
    channel: None
    amqp_queue: None
    redis_client: None
    api_gateway_uri: None

    def __init__(
        self, amqp_parameters, amqp_queue, redis_host, redis_port, api_gateway_uri
    ) -> None:
        self.amqp_parameters = amqp_parameters
        logging.info("amqp_parameters :: " + str(self.amqp_parameters))

        self.channel = get_channel(get_connection(amqp_parameters))
        logging.info("channel :: " + str(self.channel))

        self.amqp_queue = amqp_queue
        logging.info("amqp_queue :: " + str(self.amqp_queue))

        self.redis_client = get_redis_client(redis_host, redis_port)
        logging.info("redis_client :: " + str(self.redis_client))

        self.api_gateway_uri = api_gateway_uri
        logging.info("api_gateway_uri :: " + str(self.api_gateway_uri))

        self.controller = Controller(self.redis_client, self.api_gateway_uri)
        self.commands = Commands()

    def get_event_handler(self, routing_key):
        if routing_key == self.commands.CALCULATE:
            handler = self.controller.calculate
        elif routing_key == self.commands.GET:
            handler = self.controller.get

        return handler

    def generate_callback_main_queue(self, __amqp_parameters):
        def callback_queue(__channel, __method, __properties, __body):
            logging.info(__method.routing_key)
            logging.info(__properties)

            handler = self.get_event_handler(__method.routing_key)
            message = json.loads(__body)
            response = handler(message, __properties, __amqp_parameters)
            logging.info("response :: " + str(response))

            if response == True:
                logging.info("ACK")
                __channel.basic_ack(delivery_tag=__method.delivery_tag)
            else:
                logging.info("NACK")
                __channel.basic_nack(
                    delivery_tag=__method.delivery_tag, multiple=False, requeue=True
                )

            logging.info("[x] Done")

        return callback_queue

    def handle(self):
        self.channel.basic_qos(prefetch_count=1)

        callback = self.generate_callback_main_queue(self.amqp_parameters)
        self.channel.basic_consume(
            queue=self.amqp_queue, auto_ack=False, on_message_callback=callback
        )

        logging.info(" [*] Waiting for messages on queue %s" % str(self.amqp_queue))
        self.channel.start_consuming()
