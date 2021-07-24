FROM openjdk:18-slim

RUN apt update && apt upgrade -y && apt install git -y

RUN mkdir /java
WORKDIR /java

ADD . /java
