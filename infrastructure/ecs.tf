resource "aws_ecs_cluster" "ecs-cluster" {
  name               = var.cluster_name
  capacity_providers = ["FARGATE_SPOT", "FARGATE"]
  default_capacity_provider_strategy {
    capacity_provider = "FARGATE_SPOT"
  }
  # tags = var.tags
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_service_discovery_private_dns_namespace" "private-namespace" {
  name = "${var.cluster_name}.local"
  vpc  = aws_vpc.vpc_0.id
  tags = {
    Name = "${var.cluster_name}.local"
  }
  depends_on = [aws_ecs_cluster.ecs-cluster]
}

module "client-microservice" {
  source                           = "./modules/ecs-service-tg"
  aws_region                       = var.aws_region
  env                              = var.env
  docker_registry                  = var.docker_registry
  cluster_name                     = var.cluster_name
  image_version                    = var.image_version
  subnet_ids                       = aws_subnet.private_subnet.*.id
  security_group_ids               = [aws_security_group.client-microservice-sg.id]
  target_group_id                  = aws_alb_target_group.client-microservice-tg.id
  app_config                       = var.app_config_client_microservice
  app_config_container_port        = var.app_config_client_microservice_container_port
  app_config_container_environment = var.app_config_client_microservice_container_environment
  private_namespace_id             = aws_service_discovery_private_dns_namespace.private-namespace.id
  dependencies = [
    aws_subnet.private_subnet.*.id,
    aws_security_group.client-microservice-sg.id,
    aws_alb_target_group.client-microservice-tg.id,
    aws_ecs_cluster.ecs-cluster.id,
    aws_service_discovery_private_dns_namespace.private-namespace,
    rabbitmq_vhost.vhost,
    local.amqp_conn_string,
    local.mongo_conn_string,
    local.redis_conn_string,
    local.postgres_conn_string,
  ]
  tags = var.tags
}

module "users-microservice" {
  source                    = "./modules/ecs-service"
  aws_region                = var.aws_region
  env                       = var.env
  docker_registry           = var.docker_registry
  cluster_name              = var.cluster_name
  image_version             = var.image_version
  subnet_ids                = aws_subnet.private_subnet.*.id
  security_group_ids        = [aws_security_group.users-microservice-sg.id]
  app_config                = var.app_config_users_microservice
  app_config_container_port = var.app_config_users_microservice_container_port
  app_config_container_environment = concat(var.app_config_users_microservice_container_environment, [
    { "name" : "MONGO_CONN_STRING", "value" : local.mongo_conn_string },
    { "name" : "REDIS_CONN_STRING", "value" : local.redis_conn_string },
  ])
  private_namespace_id = aws_service_discovery_private_dns_namespace.private-namespace.id
  dependencies = [
    aws_subnet.private_subnet.*.id,
    aws_security_group.users-microservice-sg.id,
    aws_ecs_cluster.ecs-cluster.id,
    aws_service_discovery_private_dns_namespace.private-namespace,
    rabbitmq_vhost.vhost,
    local.amqp_conn_string,
    local.mongo_conn_string,
    local.redis_conn_string,
    local.postgres_conn_string,
  ]
  tags = var.tags
}

module "books-microservice" {
  source                    = "./modules/ecs-service"
  aws_region                = var.aws_region
  env                       = var.env
  docker_registry           = var.docker_registry
  cluster_name              = var.cluster_name
  image_version             = var.image_version
  subnet_ids                = aws_subnet.private_subnet.*.id
  security_group_ids        = [aws_security_group.books-microservice-sg.id]
  app_config                = var.app_config_books_microservice
  app_config_container_port = var.app_config_books_microservice_container_port
  app_config_container_environment = concat(var.app_config_books_microservice_container_environment, [
    { "name" : "AMQP_CONN_STRING", "value" : local.amqp_conn_string },
    { "name" : "POSTGRES_HOST", "value" : aws_db_instance.postgres.address },
    { "name" : "POSTGRES_PORT", "value" : aws_db_instance.postgres.port },
    { "name" : "POSTGRES_USER", "value" : var.postgres_username },
    { "name" : "POSTGRES_PASS", "value" : var.postgres_password },
  ])
  private_namespace_id = aws_service_discovery_private_dns_namespace.private-namespace.id
  dependencies = [
    aws_subnet.private_subnet.*.id,
    aws_security_group.books-microservice-sg.id,
    aws_ecs_cluster.ecs-cluster.id,
    aws_service_discovery_private_dns_namespace.private-namespace,
    rabbitmq_vhost.vhost,
    local.amqp_conn_string,
    local.mongo_conn_string,
    local.redis_conn_string,
    local.postgres_conn_string,
  ]
  tags = var.tags
}

module "recommendations-microservice" {
  source                    = "./modules/ecs-service"
  aws_region                = var.aws_region
  env                       = var.env
  docker_registry           = var.docker_registry
  cluster_name              = var.cluster_name
  image_version             = var.image_version
  subnet_ids                = aws_subnet.private_subnet.*.id
  security_group_ids        = [aws_security_group.recommendations-microservice-sg.id]
  app_config                = var.app_config_recommendations_microservice
  app_config_container_port = var.app_config_recommendations_microservice_container_port
  app_config_container_environment = concat(var.app_config_recommendations_microservice_container_environment, [
    { "name" : "GRPC_CONN_STRING", "value" : "${var.app_config_users_microservice.name}-${var.env}.${aws_service_discovery_private_dns_namespace.private-namespace.name}:50051" },
    { "name" : "AMQP_CONN_STRING", "value" : local.amqp_conn_string },
    { "name" : "REDIS_CONN_STRING", "value" : local.redis_conn_string },
  ])
  private_namespace_id = aws_service_discovery_private_dns_namespace.private-namespace.id
  dependencies = [
    aws_subnet.private_subnet.*.id,
    aws_security_group.recommendations-microservice-sg.id,
    aws_ecs_cluster.ecs-cluster.id,
    aws_service_discovery_private_dns_namespace.private-namespace,
    rabbitmq_vhost.vhost,
    local.amqp_conn_string,
    local.mongo_conn_string,
    local.redis_conn_string,
    local.postgres_conn_string,
  ]
  tags = var.tags
}

module "api-gateway" {
  source                    = "./modules/ecs-service-tg"
  aws_region                = var.aws_region
  env                       = var.env
  docker_registry           = var.docker_registry
  cluster_name              = var.cluster_name
  image_version             = var.image_version
  subnet_ids                = aws_subnet.private_subnet.*.id
  security_group_ids        = [aws_security_group.api-gateway-sg.id]
  target_group_id           = aws_alb_target_group.api-gateway-tg.id
  app_config                = var.app_config_api_gateway
  app_config_container_port = var.app_config_api_gateway_container_port
  app_config_container_environment = concat(var.app_config_api_gateway_container_environment, [
    { "name" : "GRPC_CONN_STRING", "value" : "${var.app_config_users_microservice.name}-${var.env}.${aws_service_discovery_private_dns_namespace.private-namespace.name}:50051" },
    { "name" : "AMQP_CONN_STRING", "value" : local.amqp_conn_string },
    { "name" : "REDIS_CONN_STRING", "value" : local.redis_conn_string },
  ])
  private_namespace_id = aws_service_discovery_private_dns_namespace.private-namespace.id
  dependencies = [
    aws_subnet.private_subnet.*.id,
    aws_security_group.api-gateway-sg.id,
    aws_alb_target_group.api-gateway-tg.id,
    aws_ecs_cluster.ecs-cluster.id,
    aws_service_discovery_private_dns_namespace.private-namespace,
    rabbitmq_vhost.vhost,
    local.amqp_conn_string,
    local.mongo_conn_string,
    local.redis_conn_string,
    local.postgres_conn_string,
  ]
  tags = var.tags
}
