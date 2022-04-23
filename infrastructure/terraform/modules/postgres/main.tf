resource "null_resource" "dependency_getter" {
  provisioner "local-exec" {
    command = "echo ${length(var.dependencies)}"
  }
}

resource "aws_db_instance" "postgres" {
  identifier             = var.postgres_identifier
  name                   = var.postgres_db_name
  engine                 = "postgres"
  engine_version         = var.postgres_engine_version
  allocated_storage      = var.postgres_allocated_storage
  apply_immediately      = true
  instance_class         = var.postgres_instance_type
  skip_final_snapshot    = true
  storage_encrypted      = false
  db_subnet_group_name   = aws_db_subnet_group.postgres-subnet-group.name
  vpc_security_group_ids = var.sg_ids
  availability_zone      = var.az_names[0]
  multi_az               = false
  publicly_accessible    = true
  port                   = 5432
  username               = var.postgres_username
  password               = var.postgres_password
  depends_on = [
    aws_db_subnet_group.postgres-subnet-group,
    null_resource.dependency_getter,
  ]
  lifecycle {
    ignore_changes = [
      availability_zone,
    ]
  }
}

resource "aws_db_subnet_group" "postgres-subnet-group" {
  name       = "postgres-subnet-group-${var.postgres_identifier}"
  subnet_ids = var.subnet_ids
  depends_on = [
    null_resource.dependency_getter,
  ]
}
