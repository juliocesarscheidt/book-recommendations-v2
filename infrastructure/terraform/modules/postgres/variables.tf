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

######################### Postgres Variables #########################
variable "postgres_identifier" {
  description = "Postgres identifier"
  type        = string
  default     = "postgres"
}

variable "postgres_db_name" {
  description = "Postgres DB name"
  type        = string
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

variable "postgres_engine_version" {
  description = "Postgres engine version"
  type        = string
  default     = "9.6"
}

variable "postgres_allocated_storage" {
  description = "Postgres allocated storage"
  type        = number
  default     = 20
}
