output "amqp_conn_string" {
  value = local.amqp_conn_string
}

output "mongo_conn_string" {
  value = local.mongo_conn_string
}

output "redis_conn_string" {
  value = local.redis_conn_string
}

# output "alb_record_fqdn" {
#   value = aws_route53_record.alb_record.fqdn
# }

output "postgres" {
  value = aws_db_instance.postgres
}
