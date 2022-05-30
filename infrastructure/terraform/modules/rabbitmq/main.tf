terraform {
  required_version = "~> 0.14.3"

  required_providers {
    rabbitmq = {
      source  = "cyrilgdn/rabbitmq"
      version = "1.6.0"
    }
  }
}

resource "null_resource" "dependency_getter" {
  provisioner "local-exec" {
    command = "echo ${length(var.dependencies)}"
  }
}

resource "aws_mq_broker" "rabbitmq" {
  broker_name         = var.rabbitmq_broker_name
  engine_type         = "RabbitMQ"
  engine_version      = var.rabbitmq_engine_version
  deployment_mode     = var.rabbitmq_deployment_mode
  publicly_accessible = true
  # security_groups     = var.sg_ids
  subnet_ids         = var.subnet_ids
  host_instance_type = var.rabbitmq_instance_type
  user {
    username = var.rabbitmq_username
    password = var.rabbitmq_password
  }
  depends_on = [
    null_resource.dependency_getter,
  ]
}

provider "rabbitmq" {
  endpoint = "https://${replace(local.amqp_address, ":5671", "")}"
  username = var.rabbitmq_username
  password = var.rabbitmq_password
}

resource "rabbitmq_vhost" "vhost" {
  name = "/"
  depends_on = [
    aws_mq_broker.rabbitmq
  ]
}

resource "rabbitmq_permissions" "rabbitmq" {
  user  = var.rabbitmq_username
  vhost = rabbitmq_vhost.vhost.name
  permissions {
    configure = ".*"
    write     = ".*"
    read      = ".*"
  }
  depends_on = [rabbitmq_vhost.vhost]
}

resource "rabbitmq_queue" "rabbitmq-books-queue" {
  name  = "books_queue"
  vhost = rabbitmq_permissions.rabbitmq.vhost
  settings {
    durable     = true
    auto_delete = false
  }
  depends_on = [rabbitmq_permissions.rabbitmq]
}

resource "rabbitmq_exchange" "rabbitmq-books-exchange" {
  name  = "books_exchange"
  vhost = rabbitmq_permissions.rabbitmq.vhost
  settings {
    type        = "topic"
    durable     = true
    auto_delete = false
  }
  depends_on = [rabbitmq_permissions.rabbitmq]
}

resource "rabbitmq_binding" "rabbitmq-books-binding" {
  source           = rabbitmq_exchange.rabbitmq-books-exchange.name
  vhost            = rabbitmq_vhost.vhost.name
  destination      = rabbitmq_queue.rabbitmq-books-queue.name
  destination_type = "queue"
  routing_key      = "books.#"
  depends_on = [
    rabbitmq_exchange.rabbitmq-books-exchange,
    rabbitmq_vhost.vhost,
    rabbitmq_queue.rabbitmq-books-queue,
  ]
}

variable "recommendations-queue-arguments" {
  default = <<EOF
{
  "x-max-priority":10
}
EOF
}

resource "rabbitmq_queue" "rabbitmq-recommendations-queue" {
  name  = "recommendations_queue"
  vhost = rabbitmq_permissions.rabbitmq.vhost
  settings {
    durable        = true
    auto_delete    = false
    arguments_json = var.recommendations-queue-arguments
  }
  depends_on = [rabbitmq_permissions.rabbitmq]
}

resource "rabbitmq_exchange" "rabbitmq-recommendations-exchange" {
  name  = "recommendations_exchange"
  vhost = rabbitmq_permissions.rabbitmq.vhost
  settings {
    type        = "topic"
    durable     = true
    auto_delete = false
  }
  depends_on = [rabbitmq_permissions.rabbitmq]
}

resource "rabbitmq_binding" "rabbitmq-recommendations-binding" {
  source           = rabbitmq_exchange.rabbitmq-recommendations-exchange.name
  vhost            = rabbitmq_vhost.vhost.name
  destination      = rabbitmq_queue.rabbitmq-recommendations-queue.name
  destination_type = "queue"
  routing_key      = "recommendations.#"
  depends_on = [
    rabbitmq_exchange.rabbitmq-recommendations-exchange,
    rabbitmq_vhost.vhost,
    rabbitmq_queue.rabbitmq-recommendations-queue,
  ]
}
