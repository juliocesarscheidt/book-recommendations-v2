class AbastractAmqpAdapter:
    def __init__(self, amqp_parameters, amqp_queue) -> None:
        pass

    def get_connection(self, amqp_parameters):
        pass

    def get_channel(self, connection):
        pass

    def channel_ack(self, channel, delivery_tag):
        pass

    def channel_nack(self, channel, delivery_tag, requeue=True):
        pass
