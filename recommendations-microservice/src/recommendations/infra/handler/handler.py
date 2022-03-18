import json
import logging

from application.commands.commands import Commands
from infra.adapter.redis_adapter import get_redis_client
from infra.adapter.amqp_adapter import get_connection, get_channel
from infra.controller.controller import Controller


class Handler(object):
    controller: Controller
    commands: Commands
    amqp_parameters: None
    channel: None
    amqp_queue: None
    redis_client: None
    grpc_conn_string: None

    def __init__(
        self, amqp_parameters, amqp_queue, redis_conn_string, grpc_conn_string
    ) -> None:
        self.amqp_parameters = amqp_parameters
        logging.info("amqp_parameters :: " + str(self.amqp_parameters))

        self.channel = get_channel(get_connection(amqp_parameters))
        logging.info("channel :: " + str(self.channel))

        self.amqp_queue = amqp_queue
        logging.info("amqp_queue :: " + str(self.amqp_queue))

        self.redis_client = get_redis_client(redis_conn_string)
        logging.info("redis_client :: " + str(self.redis_client))

        self.grpc_conn_string = grpc_conn_string
        logging.info("grpc_conn_string :: " + str(self.grpc_conn_string))

        self.controller = Controller(self.redis_client, self.grpc_conn_string)
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
