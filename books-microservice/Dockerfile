FROM juliocesarmidia/java14.0.2-maven3.8.1:latest
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /usr/src/app

COPY target/*.jar ./

CMD java $JAVA_OPTIONS -jar /usr/src/app/app.jar
