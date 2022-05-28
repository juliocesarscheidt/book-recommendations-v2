class AbstractPublisher:
    def __init__(self) -> None:
        pass

    def publish_temp_queue(self, channel, reply_queue_name, response_body):
        pass

    def publish_exchange(
        self, channel, exchange, routing_key, response_body, reply_queue_name
    ):
        pass
