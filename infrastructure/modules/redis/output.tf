output "redis_conn_string" {
  value = "${join(",", aws_elasticache_cluster.redis_cluster.cache_nodes.*.address)}:6379"
}
