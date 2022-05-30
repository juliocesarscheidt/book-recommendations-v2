terraform {
  required_version = "~> 0.14.3"

  required_providers {
    rabbitmq = {
      source  = "cyrilgdn/rabbitmq"
      version = "1.6.0"
    }
  }
}
