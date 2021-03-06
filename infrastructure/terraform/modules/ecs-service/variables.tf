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

variable "subnet_ids" {
  type        = list(string)
  description = "Subnet IDs"
}

variable "security_group_ids" {
  type        = list(string)
  description = "Security group IDs"
}

variable "app_config" {
  description = "Config for app"
  type        = map(string)
}

variable "app_config_container_port" {
  type        = number
  description = "Config for app container port"
}

variable "app_config_container_environment" {
  type        = list(any)
  description = "Config for app container environment"
}

variable "private_namespace_id" {
  type        = string
  description = "ID for private namespace"
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
