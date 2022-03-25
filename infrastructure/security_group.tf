resource "aws_security_group" "redis-sg" {
  name   = "redis-sg-${var.env}"
  vpc_id = aws_vpc.vpc_0.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 6379
    to_port   = 6379
    protocol  = "tcp"
    # cidr_blocks = [aws_vpc.vpc_0.cidr_block]
    cidr_blocks = ["0.0.0.0/0"]
  }
  depends_on = [aws_vpc.vpc_0]
}

resource "aws_security_group" "mongo-sg" {
  name   = "mongo-sg-${var.env}"
  vpc_id = aws_vpc.vpc_0.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 27017
    to_port   = 27017
    protocol  = "tcp"
    # cidr_blocks = [aws_vpc.vpc_0.cidr_block]
    cidr_blocks = ["0.0.0.0/0"]
  }
  depends_on = [aws_vpc.vpc_0]
}

resource "aws_security_group" "postgres-sg" {
  name   = "postgres-sg-${var.env}"
  vpc_id = aws_vpc.vpc_0.id
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"
    # cidr_blocks = [aws_vpc.vpc_0.cidr_block]
    cidr_blocks = ["0.0.0.0/0"]
  }
  depends_on = [aws_vpc.vpc_0]
}

# resource "aws_security_group" "rabbitmq-sg" {
#   name   = "rabbitmq-sg-${var.env}"
#   vpc_id = aws_vpc.vpc_0.id
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   ingress {
#     from_port   = 5671
#     to_port     = 5671
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   ingress {
#     from_port   = 443
#     to_port     = 443
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   lifecycle {
#     create_before_destroy = true
#   }
#   depends_on = [aws_vpc.vpc_0]
# }

resource "aws_security_group" "alb-sg" {
  vpc_id = aws_vpc.vpc_0.id
  name   = "alb-sg-${var.env}"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  lifecycle {
    create_before_destroy = true
  }
  depends_on = [aws_vpc.vpc_0]
}

resource "aws_security_group" "client-microservice-sg" {
  vpc_id = aws_vpc.vpc_0.id
  name   = "client-microservice-sg-${var.env}"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port       = var.app_config_client_microservice_container_port
    to_port         = var.app_config_client_microservice_container_port
    protocol        = "tcp"
    security_groups = [aws_security_group.alb-sg.id]
  }
  lifecycle {
    create_before_destroy = true
  }
  depends_on = [aws_security_group.alb-sg]
}

resource "aws_security_group" "users-microservice-sg" {
  vpc_id = aws_vpc.vpc_0.id
  name   = "users-microservice-sg-${var.env}"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port       = var.app_config_users_microservice_container_port
    to_port         = var.app_config_users_microservice_container_port
    protocol        = "tcp"
    security_groups = [aws_security_group.api-gateway-sg.id, aws_security_group.recommendations-microservice-sg.id]
  }
  lifecycle {
    create_before_destroy = true
  }
  depends_on = [aws_security_group.api-gateway-sg, aws_security_group.recommendations-microservice-sg]
}

resource "aws_security_group" "books-microservice-sg" {
  vpc_id = aws_vpc.vpc_0.id
  name   = "books-microservice-sg-${var.env}"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  lifecycle {
    create_before_destroy = true
  }
  depends_on = [aws_security_group.alb-sg]
}

resource "aws_security_group" "recommendations-microservice-sg" {
  vpc_id = aws_vpc.vpc_0.id
  name   = "recommendations-microservice-sg-${var.env}"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  lifecycle {
    create_before_destroy = true
  }
  depends_on = [aws_security_group.alb-sg]
}

resource "aws_security_group" "api-gateway-sg" {
  vpc_id = aws_vpc.vpc_0.id
  name   = "api-gateway-sg-${var.env}"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port       = var.app_config_api_gateway_container_port
    to_port         = var.app_config_api_gateway_container_port
    protocol        = "tcp"
    security_groups = [aws_security_group.alb-sg.id]
  }
  lifecycle {
    create_before_destroy = true
  }
  depends_on = [aws_security_group.alb-sg]
}
