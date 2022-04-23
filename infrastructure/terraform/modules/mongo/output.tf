output "mongo_conn_string" {
  value = "mongodb://${var.mongo_username}:${var.mongo_password}@${aws_docdb_cluster.mongo_cluster.endpoint}:27017/?replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false"
}
