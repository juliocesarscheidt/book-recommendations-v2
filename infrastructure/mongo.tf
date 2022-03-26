resource "aws_docdb_cluster" "mongo_cluster" {
  cluster_identifier              = var.mongo_cluster_identifier
  engine                          = "docdb"
  port                            = 27017
  availability_zones              = [data.aws_availability_zones.available_azs.names[0]]
  db_subnet_group_name            = aws_docdb_subnet_group.mongo_cluster_subnet_group.name
  db_cluster_parameter_group_name = aws_docdb_cluster_parameter_group.mongo_cluster_parameter_group.name
  master_username                 = var.mongo_username
  master_password                 = var.mongo_password
  vpc_security_group_ids          = [aws_security_group.mongo-sg.id]
  backup_retention_period         = 1
  skip_final_snapshot             = true
  deletion_protection             = false
  tags = {
    Name        = var.mongo_cluster_identifier
    Environment = var.env
  }
  depends_on = [
    aws_docdb_subnet_group.mongo_cluster_subnet_group,
    aws_docdb_cluster_parameter_group.mongo_cluster_parameter_group,
    aws_security_group.mongo-sg
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
    aws_docdb_cluster.mongo_cluster
  ]
}

resource "aws_docdb_cluster_parameter_group" "mongo_cluster_parameter_group" {
  family = "docdb4.0"
  name   = "mongo-cluster-parameter-group-${var.mongo_cluster_identifier}"
  parameter {
    name  = "tls"
    value = "disabled"
  }
}

resource "aws_docdb_subnet_group" "mongo_cluster_subnet_group" {
  name       = "mongo-cluster-subnet-group-${var.mongo_cluster_identifier}"
  subnet_ids = aws_subnet.private_subnet.*.id

  depends_on = [
    aws_subnet.private_subnet,
  ]
}
