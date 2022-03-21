#!make

MAVEN_REPO?="$$(echo ~/.m2/)/repository"

build:
	docker image build \
		-t docker.io/juliocesarmidia/books-microservice:latest .

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
		mvn install

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

# mvn package -Dmaven.test.skip=true

prepare-package:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn prepare-package

test:
	docker container run -t --rm \
		--name maven-compiler \
		-v $(MAVEN_REPO):/root/.m2/repository \
		-v "$$(pwd)":/usr/src/app \
		-w /usr/src/app \
		juliocesarmidia/java14.0.2-maven3.8.1:latest \
		mvn test

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
		mvn jacoco:prepare-agent install jacoco:report

# mvn clean jacoco:prepare-agent install jacoco:report
