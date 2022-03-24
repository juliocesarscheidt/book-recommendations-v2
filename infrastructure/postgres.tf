resource "aws_db_instance" "postgres" {
  identifier             = var.postgres_identifier
  name                   = var.postgres_identifier
  allocated_storage      = 20
  apply_immediately      = true
  instance_class         = var.postgres_instance_type
  skip_final_snapshot    = true
  storage_encrypted      = true
  db_subnet_group_name   = aws_db_subnet_group.postgres-subnet-group.name
  vpc_security_group_ids = [aws_security_group.postgres-sg.id]
  availability_zone      = data.aws_availability_zones.available_azs.names[0]
  multi_az               = false
  publicly_accessible    = true
  port                   = 5432
  username               = var.postgres_username
  password               = var.postgres_password
  engine                 = "postgres"
  engine_version         = "9.6.22-R1"

  depends_on = [
    aws_security_group.postgres-sg,
    aws_db_subnet_group.postgres-subnet-group
  ]

  lifecycle {
    ignore_changes = [
      availability_zone,
    ]
  }
}

resource "aws_db_subnet_group" "postgres-subnet-group" {
  name       = "postgres-subnet-group-${var.postgres_identifier}"
  subnet_ids = aws_subnet.public_subnet.*.id
}
