FROM ubuntu:latest

RUN apt update &&\
    apt install -y wget curl unzip make tzdata software-properties-common

RUN CHROMEDRIVER_VERSION=`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE` \
    && curl -sS -o /tmp/chromedriver_linux64.zip http://chromedriver.storage.googleapis.com/$CHROMEDRIVER_VERSION/chromedriver_linux64.zip \
    && unzip /tmp/chromedriver_linux64.zip \
    && mv chromedriver /usr/local/bin/

RUN add-apt-repository ppa:longsleep/golang-backports \
    && apt update \
    && apt install -y --no-install-recommends golang-1.17

ENV TZ=Asia/Tokyo

ENV GO111MODULE=on
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /go/src
