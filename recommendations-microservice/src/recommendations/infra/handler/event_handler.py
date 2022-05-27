import os
import json
import logging

from threading import Thread

from application.commands.commands import Commands

from infra.adapter.amqp_adapter import get_connection, get_channel
from infra.controller.event_controller import EventController
from infra.adapter.publisher import Publisher
from infra.adapter.redis_adapter import RedisAdapter
from infra.adapter.grpc_adapter import GrpcAdapter


PREFETCH_COUNT = int(os.environ.get("PREFETCH_COUNT", "1"))


class EventHandler(object):
    commands: Commands
    publisher: Publisher
    redis_adapter: RedisAdapter
    grpc_adapter: GrpcAdapter
    event_controller: EventController
    amqp_parameters: None
    channel: None
    amqp_queue: None

    def __init__(
        self, amqp_parameters, amqp_queue, redis_conn_string, grpc_conn_string
    ) -> None:
        self.amqp_parameters = amqp_parameters
        logging.info("amqp_parameters :: " + str(self.amqp_parameters))

        self.channel = get_channel(get_connection(amqp_parameters))
        logging.info("channel :: " + str(self.channel))

        self.amqp_queue = amqp_queue
        logging.info("amqp_queue :: " + str(self.amqp_queue))

        self.commands = Commands()
        self.publisher = Publisher()
        self.redis_adapter = RedisAdapter(redis_conn_string)
        self.grpc_adapter = GrpcAdapter(grpc_conn_string)
        self.event_controller = EventController(
            self.publisher, self.redis_adapter, self.grpc_adapter
        )

    def calculate_recommendation(self, __message, __properties, __amqp_parameters):
        try:
            __connection = get_connection(__amqp_parameters)
            __channel = get_channel(__connection)
            # get rates
            thread = Thread(
                target=self.event_controller.calculate_recommendation,
                args=[__message, __channel, __properties],
            )
            thread.start()
            return True
        except Exception as e:
            logging.error(e)
            return False

    def get_recommendation(self, __message, __properties, __amqp_parameters):
        try:
            __connection = get_connection(__amqp_parameters)
            __channel = get_channel(__connection)
            # get recommendation
            thread = Thread(
                target=self.event_controller.get_recommendation,
                args=[__message, __channel, __properties],
            )
            thread.start()
            return True

        except Exception as e:
            logging.error(e)
            return False

    def generate_callback_queue(self, __amqp_parameters):
        def callback_queue(__channel, __method, __properties, __body):
            logging.info(__method.routing_key)
            logging.info(__properties)

            handler_cmd = self.commands.get_command(__method.routing_key)
            handler_fn = self.commands.validate_command_handler(handler_cmd, self)
            if handler_fn is None:
                logging.info("[x] Invalid Handler")
                # TODO: send to a dead-letter queue
                self.publisher.channel_nack(__channel, __method.delivery_tag, False)
                return

            message = json.loads(__body)
            # call the method dynamically
            response = handler_fn(message, __properties, __amqp_parameters)
            logging.info("response :: " + str(response))

            if response == True:
                logging.info("ACK")
                self.publisher.channel_ack(__channel, __method.delivery_tag)
            else:
                logging.info("NACK")
                # TODO: add some exponential backoff on retries
                self.publisher.channel_nack(__channel, __method.delivery_tag, True)

            logging.info("[x] Done")

        return callback_queue

    def start(self):
        self.channel.basic_qos(prefetch_count=PREFETCH_COUNT)

        callback = self.generate_callback_queue(self.amqp_parameters)
        self.channel.basic_consume(
            queue=self.amqp_queue, auto_ack=False, on_message_callback=callback
        )

        logging.info(" [*] Waiting for messages on queue %s" % str(self.amqp_queue))
        self.channel.start_consuming()
