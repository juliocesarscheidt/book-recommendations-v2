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

variable "mongo_deletion_protection" {
  description = "Mongo deletion protection"
  type        = bool
  default     = false
}
