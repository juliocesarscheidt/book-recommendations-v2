variable "env" {
  type        = string
  description = "Environment"
  default     = "development"
}

variable "dependencies" {
  type        = any
  default     = []
  description = "Dependencies"
}

variable "tags" {
  type        = map(string)
  description = "Additional tags (_e.g._ { BusinessUnit : ABC })"
  default     = {}
}

######################### Module Variables #########################
# variable "sg_ids" {
#   type        = list(string)
#   description = "Security group IDs"
#   default     = []
# }

variable "subnet_ids" {
  type        = list(string)
  description = "Subnet IDs"
  default     = []
}

######################### RabbitMQ Variables #########################
variable "rabbitmq_broker_name" {
  description = "RabbitMQ broker name"
  type        = string
  default     = "rabbitmq"
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

variable "rabbitmq_username" {
  type        = string
  description = "RabbitMQ username"
  default     = "rabbitmq"
}

variable "rabbitmq_password" {
  type        = string
  description = "RabbitMQ password"
}
