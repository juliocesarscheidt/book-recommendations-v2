resource "aws_ecs_cluster" "ecs-cluster" {
  name               = var.cluster_name
  capacity_providers = ["FARGATE_SPOT", "FARGATE"]
  default_capacity_provider_strategy {
    capacity_provider = "FARGATE_SPOT"
  }
  # tags = var.tags
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_service_discovery_private_dns_namespace" "private-namespace" {
  name = "${var.cluster_name}.local"
  vpc  = aws_vpc.vpc_0.id
  tags = {
    Name = "${var.cluster_name}.local"
  }
  depends_on = [aws_ecs_cluster.ecs-cluster]
}
