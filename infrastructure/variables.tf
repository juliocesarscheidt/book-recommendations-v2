variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "env" {
  type        = string
  description = "Environment"
  default     = "development"
}


variable "root_domain" {
  type        = string
  description = "The root domain"
}

variable "certificate_arn" {
  type        = string
  description = "The certificate ARN"
}


variable "docker_registry" {
  type        = string
  description = "Docker registry"
}

variable "cluster_name" {
  type        = string
  description = "Cluster name"
}

variable "image_version" {
  type        = string
  description = "Image version"
}


######################### RabbitMQ Variables #########################
variable "rabbitmq_broker_name" {
  description = "RabbitMQ broker name"
  type        = string
  default     = "rabbitmq"
}

variable "rabbitmq_username" {
  type        = string
  description = "RabbitMQ username"
  default     = "rabbitmq"
}

variable "rabbitmq_password" {
  type        = string
  description = "RabbitMQ password"
}

variable "rabbitmq_engine_version" {
  type        = string
  description = "RabbitMQ engine version"
  default     = "3.8.22"
}

variable "rabbitmq_deployment_mode" {
  type        = string
  description = "RabbitMQ deployment mode"
  default     = "SINGLE_INSTANCE"
}

variable "rabbitmq_instance_type" {
  type        = string
  description = "RabbitMQ instance type"
  default     = "mq.t3.micro"
}


######################### Mongo Variables #########################
variable "mongo_cluster_identifier" {
  description = "Mongo cluster identifier"
  type        = string
  default     = "mongo"
}

variable "mongo_instance_type" {
  description = "Mongo instance type"
  type        = string
  default     = "db.t3.medium"
}

variable "mongo_username" {
  description = "Mongo username"
  type        = string
}

variable "mongo_password" {
  description = "Mongo password"
  type        = string
}


######################### Redis Variables #########################
variable "redis_cluster_name" {
  description = "Redis cluster name"
  type        = string
  default     = "redis"
}

variable "redis_instance_type" {
  description = "Redis instance type"
  type        = string
  default     = "cache.t2.small"
}

variable "redis_cache_nodes" {
  description = "Redis cache nodes"
  type        = number
  default     = 1
}

variable "redis_engine_version" {
  description = "Redis engine version"
  type        = string
  default     = "5.0.6"
}


######################### Postgres Variables #########################
variable "postgres_identifier" {
  description = "Postgres identifier"
  type        = string
  default     = "postgres"
}

variable "postgres_instance_type" {
  description = "Postgres instance type"
  type        = string
  default     = "db.t2.micro"
}

variable "postgres_username" {
  description = "Postgres username"
  type        = string
}

variable "postgres_password" {
  description = "Postgres password"
  type        = string
}


# apps config
# variable "app_config_stock_ui" {
#   description = "Config for app stock-ui"
#   type        = map(string)
#   default     = {}
# }
# variable "app_config_stock_ui_container_port" {
#   type        = number
#   description = "Config for app stock-ui container ports"
# }
# variable "app_config_stock_ui_container_environment" {
#   type        = list(any)
#   description = "Config for app stock-ui container environment"
# }


# variable "app_config_stock_api" {
#   description = "Config for app stock-api"
#   type        = map(any)
#   default     = {}
# }
# variable "app_config_stock_api_container_port" {
#   type        = number
#   description = "Config for app stock-api container ports"
# }
# variable "app_config_stock_api_container_environment" {
#   type        = list(any)
#   description = "Config for app stock-api container environment"
# }


# variable "app_config_stock_crawler" {
#   description = "Config for app stock-crawler"
#   type        = map(any)
#   default     = {}
# }
# variable "app_config_stock_crawler_container_port" {
#   type        = number
#   description = "Config for app stock-crawler container ports"
# }
# variable "app_config_stock_crawler_container_environment" {
#   type        = list(any)
#   description = "Config for app stock-crawler container environment"
# }


variable "tags" {
  type        = map(string)
  description = "Additional tags (_e.g._ { BusinessUnit : ABC })"
  default     = {}
}
