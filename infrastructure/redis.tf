resource "aws_elasticache_cluster" "redis_cluster" {
  cluster_id           = var.redis_cluster_name
  engine               = "redis"
  node_type            = var.redis_instance_type
  num_cache_nodes      = var.redis_cache_nodes
  parameter_group_name = "default.redis5.0"
  engine_version       = var.redis_engine_version
  port                 = 6379
  subnet_group_name    = aws_elasticache_subnet_group.redis-subnet-group.name
  security_group_ids   = [aws_security_group.redis-sg.id]
  availability_zone    = data.aws_availability_zones.available_azs.names[0]
  az_mode              = "single-az"

  tags = {
    Name        = var.redis_cluster_name
    Environment = var.env
  }

  depends_on = [
    aws_security_group.redis-sg,
    aws_elasticache_subnet_group.redis-subnet-group
  ]

  lifecycle {
    ignore_changes = [
      availability_zone,
    ]
  }
}

resource "aws_elasticache_subnet_group" "redis-subnet-group" {
  name       = "redis-subnet-group-${var.redis_cluster_name}"
  subnet_ids = aws_subnet.private_subnet.*.id

  depends_on = [
    aws_subnet.private_subnet,
  ]
}
