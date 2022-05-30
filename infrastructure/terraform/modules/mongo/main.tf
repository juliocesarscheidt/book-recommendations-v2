terraform {
  required_version = "~> 0.14.3"
}

resource "null_resource" "dependency_getter" {
  provisioner "local-exec" {
    command = "echo ${length(var.dependencies)}"
  }
}

resource "aws_docdb_cluster" "mongo_cluster" {
  cluster_identifier              = var.mongo_cluster_identifier
  engine                          = "docdb"
  port                            = 27017
  availability_zones              = var.az_names
  db_subnet_group_name            = aws_docdb_subnet_group.mongo_cluster_subnet_group.name
  db_cluster_parameter_group_name = aws_docdb_cluster_parameter_group.mongo_cluster_parameter_group.name
  master_username                 = var.mongo_username
  master_password                 = var.mongo_password
  vpc_security_group_ids          = var.sg_ids
  backup_retention_period         = 1
  skip_final_snapshot             = true
  deletion_protection             = var.mongo_deletion_protection
  tags = {
    Name        = var.mongo_cluster_identifier
    Environment = var.env
  }
  depends_on = [
    aws_docdb_subnet_group.mongo_cluster_subnet_group,
    aws_docdb_cluster_parameter_group.mongo_cluster_parameter_group,
    null_resource.dependency_getter,
  ]
  lifecycle {
    ignore_changes = [
      availability_zones,
    ]
  }
}

resource "aws_docdb_cluster_instance" "mongo_cluster_instances" {
  count              = 1
  identifier         = "mongo-cluster-instances-${var.mongo_cluster_identifier}-${count.index}"
  cluster_identifier = aws_docdb_cluster.mongo_cluster.id
  instance_class     = var.mongo_instance_type
  depends_on = [
    aws_docdb_cluster.mongo_cluster,
    null_resource.dependency_getter,
  ]
}

resource "aws_docdb_cluster_parameter_group" "mongo_cluster_parameter_group" {
  family = "docdb4.0"
  name   = "mongo-cluster-parameter-group-${var.mongo_cluster_identifier}"
  parameter {
    name  = "tls"
    value = "disabled"
  }
  depends_on = [
    null_resource.dependency_getter,
  ]
}

resource "aws_docdb_subnet_group" "mongo_cluster_subnet_group" {
  name       = "mongo-cluster-subnet-group-${var.mongo_cluster_identifier}"
  subnet_ids = var.subnet_ids
  depends_on = [
    null_resource.dependency_getter,
  ]
}
