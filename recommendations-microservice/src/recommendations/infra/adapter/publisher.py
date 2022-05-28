import pika

from application.adapter.abstract_publisher import AbstractPublisher


class Publisher(AbstractPublisher):
    def publish_temp_queue(self, channel, reply_queue_name, response_body):
        channel.basic_publish(
            exchange="",
            routing_key=reply_queue_name,
            body=response_body,
            properties=pika.BasicProperties(
                content_type="application/json", delivery_mode=2,
            ),
        )

    def publish_exchange(
        self, channel, exchange, routing_key, response_body, reply_queue_name
    ):
        channel.basic_publish(
            exchange=exchange,
            routing_key=routing_key,
            body=response_body,
            properties=pika.BasicProperties(
                content_type="application/json",
                reply_to=reply_queue_name,
                delivery_mode=2,
            ),
        )
