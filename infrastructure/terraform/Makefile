#!make

# general variables
ENV?=development
PLAN_FILE?=tfplan
CLUSTER_NAME?=ecs-cluster-$(ENV)
IMAGE_VERSION?=0.0.1
# AWS variables
AWS_ACCESS_KEY_ID?=
AWS_SECRET_ACCESS_KEY?=
AWS_DEFAULT_REGION?=us-east-1
# Terraform backend variables
AWS_BACKEND_BUCKET?=
AWS_BACKEND_ACCESS_KEY_ID?=
AWS_BACKEND_SECRET_ACCESS_KEY?=
AWS_BACKEND_REGION?=sa-east-1
# docker variables
DOCKER_REGISTRY?=000000000000.dkr.ecr.$(AWS_DEFAULT_REGION).amazonaws.com

docker-login:
	aws ecr get-login-password --region $(AWS_DEFAULT_REGION) | \
		docker login --username AWS $(DOCKER_REGISTRY) --password-stdin

build-image:
	docker image build --tag $(DOCKER_REGISTRY)/api-gateway:$(IMAGE_VERSION) ../../api-gateway
	docker image build --tag $(DOCKER_REGISTRY)/books-microservice:$(IMAGE_VERSION) ../../books-microservice
	docker image build --tag $(DOCKER_REGISTRY)/users-microservice:$(IMAGE_VERSION) ../../users-microservice
	docker image build --tag $(DOCKER_REGISTRY)/recommendations-microservice:$(IMAGE_VERSION) ../../recommendations-microservice
	docker image build --tag $(DOCKER_REGISTRY)/client-microservice:$(IMAGE_VERSION) ../../client-microservice

create-ecr-repo:
	aws ecr describe-repositories --repository-names api-gateway --region $(AWS_DEFAULT_REGION) || \
		aws ecr create-repository --repository-name api-gateway --region $(AWS_DEFAULT_REGION)
	aws ecr describe-repositories --repository-names books-microservice --region $(AWS_DEFAULT_REGION) || \
		aws ecr create-repository --repository-name books-microservice --region $(AWS_DEFAULT_REGION)
	aws ecr describe-repositories --repository-names users-microservice --region $(AWS_DEFAULT_REGION) || \
		aws ecr create-repository --repository-name users-microservice --region $(AWS_DEFAULT_REGION)
	aws ecr describe-repositories --repository-names recommendations-microservice --region $(AWS_DEFAULT_REGION) || \
		aws ecr create-repository --repository-name recommendations-microservice --region $(AWS_DEFAULT_REGION)
	aws ecr describe-repositories --repository-names client-microservice --region $(AWS_DEFAULT_REGION) || \
		aws ecr create-repository --repository-name client-microservice --region $(AWS_DEFAULT_REGION)

push-image: docker-login build-image create-ecr-repo
	docker image push $(DOCKER_REGISTRY)/api-gateway:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/books-microservice:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/users-microservice:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/recommendations-microservice:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/client-microservice:$(IMAGE_VERSION)

create-backend-bucket:
	aws s3 ls s3://$(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION) || \
		aws s3api create-bucket --bucket $(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION) \
			--create-bucket-configuration LocationConstraint=$(AWS_BACKEND_REGION) --acl private

create-workspace:
	-@terraform workspace new development 2> /dev/null
	-@terraform workspace new production 2> /dev/null
	terraform workspace select $(ENV)

init: create-backend-bucket
	terraform init -upgrade=true \
		-backend-config="bucket=$(AWS_BACKEND_BUCKET)" \
		-backend-config="key=state.tfstate" \
		-backend-config="region=$(AWS_BACKEND_REGION)" \
		-backend-config="workspace_key_prefix=terraform/ecs-cluster" \
		-backend-config="access_key=$(AWS_BACKEND_ACCESS_KEY_ID)" \
		-backend-config="secret_key=$(AWS_BACKEND_SECRET_ACCESS_KEY)" \
		-backend-config="encrypt=true"
	make create-workspace
	make plan

fmt:
	terraform fmt -write=true -recursive

validate:
	terraform validate

lint:
	docker container run --rm -t \
		--name tflint \
		--env TFLINT_LOG=debug \
		-v $$(pwd):/data \
		ghcr.io/terraform-linters/tflint

refresh:
	terraform refresh	-var-file=$(ENV).tfvars

output:
	terraform output

plan: validate lint
	terraform plan \
		-out=$(PLAN_FILE) \
		-var-file=$(ENV).tfvars \
		-var aws_region="$(AWS_DEFAULT_REGION)" \
		-var docker_registry="$(DOCKER_REGISTRY)" \
		-var cluster_name="$(CLUSTER_NAME)" \
		-var image_version="$(IMAGE_VERSION)" \
		-input=false

apply:
	terraform apply $(PLAN_FILE)
	@echo "Access the application on https://$$(terraform output app_dns)"

destroy:
	terraform destroy \
		-var-file=$(ENV).tfvars \
		-auto-approve
