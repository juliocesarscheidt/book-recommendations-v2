import pika

from application.adapter.abstract_amqp_adapter import AbastractAmqpAdapter


class AmqpAdapter(AbastractAmqpAdapter):
    amqp_parameters: None
    amqp_queue: None
    connection: None
    channel: None

    def __init__(self, amqp_parameters, amqp_queue) -> None:
        self.amqp_parameters = amqp_parameters
        self.amqp_queue = amqp_queue
        self.connection = self.get_connection(self.amqp_parameters)
        self.channel = self.get_channel(self.connection)

    def get_connection(self, amqp_parameters):
        return pika.BlockingConnection(amqp_parameters)

    def get_channel(self, connection):
        return connection.channel()

    def channel_ack(self, channel, delivery_tag):
        channel.basic_ack(delivery_tag=delivery_tag)

    def channel_nack(self, channel, delivery_tag, requeue=True):
        channel.basic_nack(delivery_tag=delivery_tag, multiple=False, requeue=requeue)
