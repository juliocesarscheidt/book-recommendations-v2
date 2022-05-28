import os
import json

from threading import Thread

from application.commands.commands import Commands
from infra.factory.event_controller_factory import EventControllerFactory
from infra.controller.event_controller import EventController
from infra.adapter.amqp_adapter import AmqpAdapter
from infra.common.logger import logger


PREFETCH_COUNT = int(os.environ.get("PREFETCH_COUNT", "1"))


class EventHandler:
    commands: Commands
    amqp_adapter: AmqpAdapter
    event_controller: EventController

    def __init__(
        self, amqp_parameters, amqp_queue, redis_conn_string, grpc_conn_string
    ) -> None:
        self.commands = Commands()
        self.amqp_adapter = AmqpAdapter(amqp_parameters, amqp_queue)
        event_controller_factory = EventControllerFactory(
            redis_conn_string, grpc_conn_string
        )
        self.event_controller = event_controller_factory.create_event_controller()

    def calculate_recommendation(self, __message, __properties, __amqp_parameters):
        try:
            __connection = self.amqp_adapter.get_connection(__amqp_parameters)
            __channel = self.amqp_adapter.get_channel(__connection)
            # get rates
            thread = Thread(
                target=self.event_controller.calculate_recommendation,
                args=[__message, __channel, __properties],
            )
            thread.start()
            return True
        except Exception as e:
            logger.error(e)
            return False

    def get_recommendation(self, __message, __properties, __amqp_parameters):
        try:
            __connection = self.amqp_adapter.get_connection(__amqp_parameters)
            __channel = self.amqp_adapter.get_channel(__connection)
            # get recommendation
            thread = Thread(
                target=self.event_controller.get_recommendation,
                args=[__message, __channel, __properties],
            )
            thread.start()
            return True

        except Exception as e:
            logger.error(e)
            return False

    def generate_callback_queue(self, __amqp_parameters):
        def callback_queue(__channel, __method, __properties, __body):
            logger.info(__method.routing_key)
            logger.info(__properties)

            handler_cmd = self.commands.get_command(__method.routing_key)
            handler_fn = self.commands.validate_command_handler(handler_cmd, self)
            if handler_fn is None:
                logger.info("[x] Invalid Handler")
                # TODO: send to a dead-letter queue
                self.amqp_adapter.channel_nack(__channel, __method.delivery_tag, False)
                return

            message = json.loads(__body)
            # call the method dynamically
            response = handler_fn(message, __properties, __amqp_parameters)
            logger.info("response :: " + str(response))

            if response == True:
                logger.info("ACK")
                self.amqp_adapter.channel_ack(__channel, __method.delivery_tag)
            else:
                logger.info("NACK")
                # TODO: add some exponential backoff on retries
                self.amqp_adapter.channel_nack(__channel, __method.delivery_tag, True)

            logger.info("[x] Done")

        return callback_queue

    def start(self):
        self.amqp_adapter.channel.basic_qos(prefetch_count=PREFETCH_COUNT)

        callback = self.generate_callback_queue(self.amqp_adapter.amqp_parameters)
        self.amqp_adapter.channel.basic_consume(
            queue=self.amqp_adapter.amqp_queue,
            auto_ack=False,
            on_message_callback=callback,
        )

        logger.info(
            " [*] Waiting for messages on queue %s" % str(self.amqp_adapter.amqp_queue)
        )
        self.amqp_adapter.channel.start_consuming()
