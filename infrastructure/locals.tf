locals {
  subnets_offset = min(length(data.aws_availability_zones.available_azs.names), 3)
  public_subnets = [
    for index in range(0, local.subnets_offset) : cidrsubnet(aws_vpc.vpc_0.cidr_block, 8, index)
  ]
  private_subnets = [
    for index in range(0, local.subnets_offset) : cidrsubnet(aws_vpc.vpc_0.cidr_block, 8, index + local.subnets_offset)
  ]
  # amqp
  amqp_address     = replace(aws_mq_broker.rabbitmq.instances.*.endpoints[0][0], "amqps://", "")
  amqp_conn_string = "amqps://${var.rabbitmq_username}:${var.rabbitmq_password}@${local.amqp_address}/"
  # mongo
  mongo_conn_string = "mongodb://${var.mongo_username}:${var.mongo_password}@${aws_docdb_cluster.mongo_cluster.endpoint}:27017/?replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false"
  # redis
  redis_address     = aws_elasticache_cluster.redis_cluster.cache_nodes.*.address
  redis_conn_string = "${join(",", local.redis_address)}:6379"
  # postgres
  postgres_conn_string = "${aws_db_instance.postgres.address}:${aws_db_instance.postgres.port}"
}
