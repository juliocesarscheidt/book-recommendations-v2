#!make

build:
	docker image build \
		--tag api-gateway:latest \
		--file development.Dockerfile .

lint-app:
	docker container run -t --rm \
		--name api-gateway-lint \
		-v "$$(pwd)":/go/src/app \
		-w /go/src/app \
		api-gateway:latest go vet .

test-app:
	@echo "no test"
