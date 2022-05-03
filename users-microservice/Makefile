#!make

build:
	docker image build \
		--tag users-microservice:latest \
		--file development.Dockerfile .

lint-app:
	docker container run -t --rm \
		--name users-microservice-lint \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		users-microservice:latest sh -c "yarn install && yarn lint"

test-app:
	docker container run -t --rm \
		--name users-microservice-test \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		users-microservice:latest sh -c "yarn install && yarn test"
