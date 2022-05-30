locals {
  subnets_offset  = min(length(data.aws_availability_zones.available_azs.names), 3)
  public_subnets  = [for index in range(0, local.subnets_offset) : cidrsubnet(aws_vpc.vpc_0.cidr_block, 8, index)]
  private_subnets = [for index in range(0, local.subnets_offset) : cidrsubnet(aws_vpc.vpc_0.cidr_block, 8, index + local.subnets_offset)]
  # dns
  app_dns = "books-${var.env}.${var.root_domain}"
  # gRPC service
  grpc_conn_string = "${var.app_config_users_microservice.name}-${var.env}.${aws_service_discovery_private_dns_namespace.private-namespace.name}:50051"
  # amqp
  amqp_conn_string = module.rabbitmq.amqp_conn_string
  # mongo
  mongo_conn_string = module.mongo.mongo_conn_string
  # redis
  redis_conn_string = module.redis.redis_conn_string
  # postgres
  postgres_conn_string = module.postgres.postgres_conn_string
  # s3
  api_gw_bucket_name = "${var.app_config_api_gateway_bucket_name}-${var.env}"
}
