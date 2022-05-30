# this is a sample file
env             = "development"
aws_region      = "us-east-1"
root_domain     = "domain.com.br"
certificate_arn = "arn:aws:acm:us-east-1:000000000000:certificate/aaaaaaaa-aaaa-0000-0000-aaaaaaaaaaaa"
docker_registry = "000000000000.dkr.ecr.us-east-1.amazonaws.com"
cluster_name    = "ecs-cluster-development"
image_version   = "0.0.1"
# rabbitmq
rabbitmq_broker_name   = "rabbitmq-development"
rabbitmq_username      = "rabbitmquser"
rabbitmq_password      = ""
rabbitmq_instance_type = "mq.t3.micro"
# mongo
mongo_cluster_identifier = "mongo-development"
mongo_username           = "mongouser"
mongo_password           = ""
mongo_instance_type      = "db.t3.medium"
# redis
redis_cluster_name  = "redis-development"
redis_instance_type = "cache.t3.micro"
# postgres
postgres_identifier    = "postgres-development"
postgres_username      = "postgresuser"
postgres_password      = ""
postgres_instance_type = "db.t2.micro"
# e.g. 000000000000.dkr.ecr.us-east-1.amazonaws.com/client-microservice:0.0.1
app_config_client_microservice = {
  "name"                    = "client-microservice"
  "container_name"          = "client-microservice"
  "execution_role_arn"      = "arn:aws:iam::000000000000:role/AmazonECSTaskExecutionRole"
  "task_role_arn"           = ""
  "cpu"                     = "256"
  "memory"                  = "512"
  "memory_reservation"      = "256"
  "desired_count"           = "1"
  "minimum_healthy_percent" = "0"
  "maximum_percent"         = "100"
}
app_config_client_microservice_container_port = 80
app_config_client_microservice_container_environment = [
  { "name" : "VUE_APP_API_GATEWAY_CONN_STRING", "value" : "" },
]
# e.g. 000000000000.dkr.ecr.us-east-1.amazonaws.com/users-microservice:0.0.1
app_config_users_microservice = {
  "name"                    = "users-microservice"
  "container_name"          = "users-microservice"
  "execution_role_arn"      = "arn:aws:iam::000000000000:role/AmazonECSTaskExecutionRole"
  "task_role_arn"           = ""
  "cpu"                     = "256"
  "memory"                  = "512"
  "memory_reservation"      = "256"
  "desired_count"           = "1"
  "minimum_healthy_percent" = "0"
  "maximum_percent"         = "100"
}
app_config_users_microservice_container_port = 50051
app_config_users_microservice_container_environment = [
  { "name" : "NODE_ENV", "value" : "production" },
  { "name" : "MONGO_DATABASE", "value" : "user_db" },
]
# e.g. 000000000000.dkr.ecr.us-east-1.amazonaws.com/books-microservice:0.0.1
app_config_books_microservice = {
  "name"                    = "books-microservice"
  "container_name"          = "books-microservice"
  "execution_role_arn"      = "arn:aws:iam::000000000000:role/AmazonECSTaskExecutionRole"
  "task_role_arn"           = ""
  "cpu"                     = "512"
  "memory"                  = "1024"
  "memory_reservation"      = "512"
  "desired_count"           = "1"
  "minimum_healthy_percent" = "0"
  "maximum_percent"         = "100"
}
app_config_books_microservice_container_port = null
app_config_books_microservice_container_environment = [
  { "name" : "POSTGRES_DB", "value" : "book_db" },
]
# e.g. 000000000000.dkr.ecr.us-east-1.amazonaws.com/recommendations-microservice:0.0.1
app_config_recommendations_microservice = {
  "name"                    = "recommendations-microservice"
  "container_name"          = "recommendations-microservice"
  "execution_role_arn"      = "arn:aws:iam::000000000000:role/AmazonECSTaskExecutionRole"
  "task_role_arn"           = ""
  "cpu"                     = "256"
  "memory"                  = "512"
  "memory_reservation"      = "256"
  "desired_count"           = "1"
  "minimum_healthy_percent" = "0"
  "maximum_percent"         = "100"
}
app_config_recommendations_microservice_container_port = null
app_config_recommendations_microservice_container_environment = [
]
# e.g. 000000000000.dkr.ecr.us-east-1.amazonaws.com/api-gateway:0.0.1
app_config_api_gateway = {
  "name"                    = "api-gateway"
  "container_name"          = "api-gateway"
  "execution_role_arn"      = "arn:aws:iam::000000000000:role/AmazonECSTaskExecutionRole"
  "task_role_arn"           = ""
  "cpu"                     = "256"
  "memory"                  = "512"
  "memory_reservation"      = "256"
  "desired_count"           = "1"
  "minimum_healthy_percent" = "0"
  "maximum_percent"         = "100"
}
app_config_api_gateway_container_port = 3080
app_config_api_gateway_container_environment = [
]
app_config_api_gateway_bucket_name = "book-recommendations-v2-api-gw-bucket"
