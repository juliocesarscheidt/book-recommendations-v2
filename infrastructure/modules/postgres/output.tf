output "postgres_conn_string" {
  value = "${aws_db_instance.postgres.address}:${aws_db_instance.postgres.port}"
}

output "postgres_host" {
  value = aws_db_instance.postgres.address
}

output "postgres_port" {
  value = aws_db_instance.postgres.port
}
