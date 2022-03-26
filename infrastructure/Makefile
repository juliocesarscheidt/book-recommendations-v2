#!make

ENV?=development
PLAN_FILE?=tfplan
CLUSTER_NAME?=ecs-cluster-$(ENV)
IMAGE_VERSION?=0.0.1

AWS_ACCESS_KEY_ID?=
AWS_SECRET_ACCESS_KEY?=
AWS_DEFAULT_REGION?=us-east-1

DOCKER_REGISTRY?=000000000000.dkr.ecr.$(AWS_DEFAULT_REGION).amazonaws.com

AWS_BACKEND_BUCKET?=blackdevs-aws
AWS_BACKEND_REGION?=sa-east-1
AWS_BACKEND_ACCESS_KEY_ID?=
AWS_BACKEND_SECRET_ACCESS_KEY?=

create-bucket:
	aws s3 ls s3://$(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION) || \
		aws s3api create-bucket --bucket $(AWS_BACKEND_BUCKET) --region $(AWS_BACKEND_REGION) \
			--create-bucket-configuration LocationConstraint=$(AWS_BACKEND_REGION) --acl private

create-workspace:
	-@terraform workspace new development 2> /dev/null
	-@terraform workspace new staging 2> /dev/null
	-@terraform workspace new production 2> /dev/null
	terraform workspace select $(ENV)

init: create-bucket
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

refresh:
	terraform refresh	-var-file=$(ENV).tfvars

output:
	terraform output

plan: validate
	terraform plan \
		-out=$(PLAN_FILE) \
		-var-file=$(ENV).tfvars \
		-var aws_region="$(AWS_DEFAULT_REGION)" \
		-var docker_registry="$(DOCKER_REGISTRY)" \
		-var cluster_name="$(CLUSTER_NAME)" \
		-var image_version="$(IMAGE_VERSION)" \
		-input=false

deploy:
	terraform apply $(PLAN_FILE)

destroy:
	terraform destroy \
		-var-file=$(ENV).tfvars \
		-auto-approve
