#!make

build:
	docker image build \
		--tag recommendations-microservice:latest \
		--file Dockerfile .

lint-app:
	docker container run --rm \
		--name recommendations-microservice-lint \
		-v $$(pwd):/usr/src/app \
		-w /usr/src/app \
		--entrypoint "" \
		cytopia/black:latest-0.2 sh -c "black --check -v ."

test-app:
	docker container run -t --rm \
		--name recommendations-microservice-test \
		-v "$$(pwd):/usr/src/app" \
		-w /usr/src/app \
		--entrypoint "" \
		juliocesarmidia/recommendations-microservice:latest pytest -v \
			--log-level=DEBUG --full-trace ./src/test/

coverage:
	docker container run -t --rm \
		--name py-test \
		-v "$$(pwd):/usr/src/app" \
		-w /usr/src/app \
		--entrypoint "" \
		juliocesarmidia/recommendations-microservice:latest pytest -v \
			--log-level=DEBUG --full-trace \
			--cov-report term-missing \
			--cov=./src/  \
			./src/test/

report:
	docker container run -t --rm \
		--name py-test \
		-v "$$(pwd):/usr/src/app" \
		-w /usr/src/app \
		--entrypoint "" \
		juliocesarmidia/recommendations-microservice:latest pytest -v \
			--log-level=DEBUG --full-trace \
			-o junit_family=xunit1 \
			--cov-report xml:src/coverage-reports/coverage.xml \
			--junitxml=src/coverage-reports/results.xml \
			--cov=./src/ ./src/test/
