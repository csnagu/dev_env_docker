FROM golang:1.16

RUN apt update && apt upgrade -y && apt install git -y

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app
