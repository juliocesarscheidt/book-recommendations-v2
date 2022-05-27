class AbstractPublisher(object):
    def __init__(self) -> None:
        pass

    def publish_temp_queue(self, channel, reply_queue_name, response_body):
        pass

    def publish_exchange(
        self, channel, exchange, routing_key, response_body, reply_queue_name
    ):
        pass

    def channel_ack(self, channel, delivery_tag):
        pass

    def channel_nack(self, channel, delivery_tag, requeue=True):
        pass
