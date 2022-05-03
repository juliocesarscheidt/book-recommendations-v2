#!make

MAVEN_REPO?="$$(echo ~/.m2/)/repository"

build: install package
	docker image build \
		--tag books-microservice:latest \
		--file Dockerfile .

lint-app:
	@echo "no lint"

test-app:
	docker container run -t --rm \
		--name books-microservice-test \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn test

clean:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn clean

install:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn install -Dmaven.test.skip=true

compile:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn compile

package:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn package -Dmaven.test.skip=true

prepare-package:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn prepare-package

verify:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn verify

report:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn clean jacoco:prepare-agent install jacoco:report
