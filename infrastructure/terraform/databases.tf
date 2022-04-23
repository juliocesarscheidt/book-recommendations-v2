module "mongo" {
  source                   = "./modules/mongo"
  env                      = var.env
  tags                     = var.tags
  az_names                 = [data.aws_availability_zones.available_azs.names[0]]
  sg_ids                   = [aws_security_group.mongo-sg.id]
  subnet_ids               = aws_subnet.private_subnet.*.id
  mongo_cluster_identifier = var.mongo_cluster_identifier
  mongo_username           = var.mongo_username
  mongo_password           = var.mongo_password
  mongo_instance_type      = var.mongo_instance_type
  dependencies = [
    aws_security_group.mongo-sg,
    aws_subnet.private_subnet,
  ]
}

module "postgres" {
  source                 = "./modules/postgres"
  env                    = var.env
  tags                   = var.tags
  az_names               = [data.aws_availability_zones.available_azs.names[0]]
  sg_ids                 = [aws_security_group.postgres-sg.id]
  subnet_ids             = aws_subnet.private_subnet.*.id
  postgres_identifier    = var.postgres_identifier
  postgres_db_name       = "book_db"
  postgres_instance_type = var.postgres_instance_type
  postgres_username      = var.postgres_username
  postgres_password      = var.postgres_password
  dependencies = [
    aws_security_group.postgres-sg,
    aws_subnet.private_subnet,
  ]
}

module "redis" {
  source               = "./modules/redis"
  env                  = var.env
  tags                 = var.tags
  az_names             = [data.aws_availability_zones.available_azs.names[0]]
  sg_ids               = [aws_security_group.redis-sg.id]
  subnet_ids           = aws_subnet.private_subnet.*.id
  redis_cluster_name   = var.redis_cluster_name
  redis_instance_type  = var.redis_instance_type
  redis_cache_nodes    = var.redis_cache_nodes
  redis_engine_version = var.redis_engine_version
  dependencies = [
    aws_security_group.redis-sg,
    aws_subnet.private_subnet,
  ]
}
