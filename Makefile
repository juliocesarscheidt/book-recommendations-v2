#!make

ENV?=development
PLAN_FILE?=tfplan
CLUSTER_NAME?=ecs-cluster-$(ENV)
IMAGE_VERSION?=0.0.1

AWS_ACCESS_KEY_ID?=
AWS_SECRET_ACCESS_KEY?=
AWS_DEFAULT_REGION?=us-east-1

DOCKER_REGISTRY?=829560024531.dkr.ecr.$(AWS_DEFAULT_REGION).amazonaws.com

TERRAFORM_PATH?=$$(pwd)/infrastructure

docker-login:
	aws ecr get-login-password --region $(AWS_DEFAULT_REGION) | \
		docker login --username AWS $(DOCKER_REGISTRY) --password-stdin

build-image:
	docker image build --tag $(DOCKER_REGISTRY)/api-gateway:$(IMAGE_VERSION) $$(pwd)/api-gateway
	docker image build --tag $(DOCKER_REGISTRY)/books-microservice:$(IMAGE_VERSION) $$(pwd)/books-microservice
	docker image build --tag $(DOCKER_REGISTRY)/users-microservice:$(IMAGE_VERSION) $$(pwd)/users-microservice
	docker image build --tag $(DOCKER_REGISTRY)/recommendations-microservice:$(IMAGE_VERSION) $$(pwd)/recommendations-microservice
	docker image build --tag $(DOCKER_REGISTRY)/client-microservice:$(IMAGE_VERSION) $$(pwd)/client-microservice

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

init:
	cd $(TERRAFORM_PATH) && make init

deploy:
	cd $(TERRAFORM_PATH) && make deploy

destroy:
	cd $(TERRAFORM_PATH) && make destroy
