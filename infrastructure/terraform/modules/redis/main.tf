resource "null_resource" "dependency_getter" {
  provisioner "local-exec" {
    command = "echo ${length(var.dependencies)}"
  }
}

resource "aws_elasticache_cluster" "redis_cluster" {
  cluster_id           = var.redis_cluster_name
  engine               = "redis"
  node_type            = var.redis_instance_type
  num_cache_nodes      = var.redis_cache_nodes
  parameter_group_name = "default.redis5.0"
  engine_version       = var.redis_engine_version
  port                 = 6379
  az_mode              = var.redis_az_mode
  availability_zone    = var.az_names[0]
  subnet_group_name    = aws_elasticache_subnet_group.redis-subnet-group.name
  security_group_ids   = var.sg_ids
  tags = {
    Name        = var.redis_cluster_name
    Environment = var.env
  }
  depends_on = [
    aws_elasticache_subnet_group.redis-subnet-group,
    null_resource.dependency_getter,
  ]
  lifecycle {
    ignore_changes = [
      availability_zone,
    ]
  }
}

resource "aws_elasticache_subnet_group" "redis-subnet-group" {
  name       = "redis-subnet-group-${var.redis_cluster_name}"
  subnet_ids = var.subnet_ids
  depends_on = [
    null_resource.dependency_getter,
  ]
}
