module "rabbitmq" {
  source               = "./modules/rabbitmq"
  env                  = var.env
  tags                 = var.tags
  # sg_ids               = [aws_security_group.rabbitmq-sg.id]
  subnet_ids           = var.rabbitmq_deployment_mode == "SINGLE_INSTANCE" ? [aws_subnet.private_subnet.0.id] : aws_subnet.private_subnet.*.id
  rabbitmq_broker_name = var.rabbitmq_broker_name
  rabbitmq_engine_version = var.rabbitmq_engine_version
  rabbitmq_deployment_mode = var.rabbitmq_deployment_mode
  rabbitmq_instance_type = var.rabbitmq_instance_type
  rabbitmq_username = var.rabbitmq_username
  rabbitmq_password = var.rabbitmq_password
  dependencies = [
    # aws_security_group.rabbitmq-sg,
    aws_subnet.private_subnet,
  ]
}
