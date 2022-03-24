######## application load balancer ########
resource "aws_lb" "application-lb" {
  load_balancer_type         = "application"
  name                       = "application-lb"
  internal                   = false
  enable_deletion_protection = false
  idle_timeout               = 300
  subnets                    = aws_subnet.public_subnet.*.id
  security_groups            = [aws_security_group.alb-sg.id]
  tags = {
    Name = "application-lb"
  }
  lifecycle {
    create_before_destroy = true
  }
}


# ######## target groups ########
# resource "aws_alb_target_group" "client-microservice-tg" {
#   name                          = "client-microservice-tg"
#   port                          = var.app_config_client_microservice_container_port
#   protocol                      = "HTTP"
#   vpc_id                        = aws_vpc.vpc_0.id
#   load_balancing_algorithm_type = "least_outstanding_requests"
#   deregistration_delay          = 60
#   target_type                   = "ip"
#   health_check {
#     healthy_threshold   = 2
#     unhealthy_threshold = 10
#     timeout             = 30
#     interval            = 60
#     matcher             = "200-299"
#     path                = "/"
#     protocol            = "HTTP"
#     port                = var.app_config_client_microservice_container_port
#   }
#   lifecycle {
#     create_before_destroy = true
#   }
# }

# resource "aws_alb_target_group" "api-gateway-tg" {
#   name                          = "api-gateway-tg"
#   port                          = var.app_config_api_gateway_container_port
#   protocol                      = "HTTP"
#   vpc_id                        = aws_vpc.vpc_0.id
#   load_balancing_algorithm_type = "least_outstanding_requests"
#   deregistration_delay          = 60
#   target_type                   = "ip"
#   health_check {
#     healthy_threshold   = 2
#     unhealthy_threshold = 10
#     timeout             = 30
#     interval            = 60
#     matcher             = "200-299"
#     path                = "/v1/healthcheck"
#     protocol            = "HTTP"
#     port                = var.app_config_api_gateway_container_port
#   }
#   lifecycle {
#     create_before_destroy = true
#   }
# }


# ######## listeners ########
# resource "aws_alb_listener" "application-lb-listener-http" {
#   load_balancer_arn = aws_lb.application-lb.arn
#   port              = 80
#   protocol          = "HTTP"
#   default_action {
#     type = "redirect"
#     redirect {
#       port        = 443
#       protocol    = "HTTPS"
#       status_code = "HTTP_301"
#     }
#   }
# }

# resource "aws_alb_listener" "application-lb-listener-https" {
#   load_balancer_arn = aws_lb.application-lb.arn
#   port              = 443
#   protocol          = "HTTPS"
#   # ARN for SSL certificate
#   certificate_arn = var.certificate_arn
#   default_action {
#     type             = "forward"
#     target_group_arn = aws_alb_target_group.client-microservice-tg.id
#   }
# }

# ######## listener rules ########
# resource "aws_alb_listener_rule" "application-lb-listener-rule" {
#   listener_arn = aws_alb_listener.application-lb-listener-https.arn
#   priority     = 100
#   action {
#     type             = "forward"
#     target_group_arn = aws_alb_target_group.api-gateway-tg.id
#   }
#   condition {
#     path_pattern {
#       values = ["/v1/*"]
#     }
#   }
#   depends_on = [
#     aws_alb_listener.application-lb-listener-https,
#     aws_alb_target_group.api-gateway-tg,
#   ]
# }
