#!make

ENV?=development
PLAN_FILE?=tfplan
CLUSTER_NAME?=ecs-cluster-development
# DOCKER_REGISTRY?=
IMAGE_VERSION?=0.0.1

AWS_DEFAULT_REGION?=us-east-1
DOCKER_REGISTRY?=829560024531.dkr.ecr.$(AWS_DEFAULT_REGION).amazonaws.com
AWS_BACKEND_BUCKET?=blackdevs-aws

docker-login:
	-@echo "Docker login"
	aws ecr get-login-password --region $(AWS_DEFAULT_REGION) | \
		docker login --username AWS $(DOCKER_REGISTRY) --password-stdin

docker-build-image:
	-@echo "Building images"
	docker image build --tag $(DOCKER_REGISTRY)/api-gateway:$(IMAGE_VERSION) ../api-gateway
	docker image build --tag $(DOCKER_REGISTRY)/books-microservice:$(IMAGE_VERSION) ../books-microservice
	docker image build --tag $(DOCKER_REGISTRY)/users-microservice:$(IMAGE_VERSION) ../users-microservice
	docker image build --tag $(DOCKER_REGISTRY)/recommendations-microservice:$(IMAGE_VERSION) ../recommendations-microservice
	docker image build --tag $(DOCKER_REGISTRY)/client-microservice:$(IMAGE_VERSION) ../client-microservice

docker-create-repo:
	-@echo "Creating repositories"
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

push-image: docker-login docker-build-image docker-create-repo
	-@echo "Pushing images"
	docker image push $(DOCKER_REGISTRY)/api-gateway:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/books-microservice:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/users-microservice:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/recommendations-microservice:$(IMAGE_VERSION)
	docker image push $(DOCKER_REGISTRY)/client-microservice:$(IMAGE_VERSION)

create-bucket:
	-@echo "Creating bucket"
	aws s3 ls s3://$(AWS_BACKEND_BUCKET) --region $(AWS_DEFAULT_REGION) || \
		aws s3api create-bucket --bucket $(AWS_BACKEND_BUCKET) --region $(AWS_DEFAULT_REGION) \
			--create-bucket-configuration LocationConstraint=$(AWS_DEFAULT_REGION) --acl private

# init: create-bucket
init:
	-@echo "Init"
	terraform init -upgrade=true \
		-backend-config="bucket=$(AWS_BACKEND_BUCKET)" \
		-backend-config="key=state.tfstate" \
		-backend-config="region=$(AWS_DEFAULT_REGION)" \
		-backend-config="workspace_key_prefix=terraform/$(CLUSTER_NAME)" \
		-backend-config="access_key=$(AWS_ACCESS_KEY_ID)" \
		-backend-config="secret_key=$(AWS_SECRET_ACCESS_KEY)" \
		-backend-config="encrypt=true"

	-@terraform workspace new development 2> /dev/null
	-@terraform workspace new production 2> /dev/null
	terraform workspace select $(ENV)

	make plan

fmt:
	terraform fmt -write=true -recursive

validate:
	terraform validate

refresh:
	terraform refresh	-var-file=$(ENV).tfvars

output:
	terraform output

plan: validate
	-@echo "Plan"
	terraform plan \
		-out=$(PLAN_FILE) \
		-var-file=$(ENV).tfvars \
		-var aws_region="$(AWS_DEFAULT_REGION)" \
		-var docker_registry="$(DOCKER_REGISTRY)" \
		-var cluster_name="$(CLUSTER_NAME)" \
		-var image_version="$(IMAGE_VERSION)" \
		-input=false

apply:
	-@echo "Apply"
	terraform apply $(PLAN_FILE)
	terraform output -json

destroy:
	-@echo "Destroy"
	terraform destroy \
		-var-file=$(ENV).tfvars \
		-auto-approve
