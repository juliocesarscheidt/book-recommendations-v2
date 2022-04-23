variable "env" {
  type        = string
  description = "Environment"
  default     = "development"
}

variable dependencies {
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
variable "az_names" {
  type        = list(string)
  description = "Availability zonbes"
  default     = []
}

variable "sg_ids" {
  type        = list(string)
  description = "Security group IDs"
  default     = []
}

variable "subnet_ids" {
  type        = list(string)
  description = "Subnet IDs"
  default     = []
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
  default     = "cache.t3.micro"
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

variable "redis_az_mode" {
  description = "Redis AZ mode"
  type        = string
  default     = "single-az"
}
