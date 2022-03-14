import pika


def get_connection(parameters):
    return pika.BlockingConnection(parameters)


def get_channel(connection):
    return connection.channel()
