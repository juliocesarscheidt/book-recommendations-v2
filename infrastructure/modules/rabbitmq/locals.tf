locals {
  amqp_address     = replace(aws_mq_broker.rabbitmq.instances.*.endpoints[0][0], "amqps://", "")
  amqp_conn_string = "amqps://${var.rabbitmq_username}:${var.rabbitmq_password}@${local.amqp_address}/"
}
