FROM ubuntu:20.04

ENV TZ=Asia/Tokyo
ENV GO111MODULE=on
ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin
ENV ARCH amd64
ENV GOVERSION 1.17

RUN apt update \
    && apt install -y wget curl unzip make tzdata software-properties-common ca-certificates gcc libfontconfig libxkbcommon0 libatspi2.0-0 libgbm1 libnss3 libgconf-2-4 fonts-ipafont fonts-ipaexfont \
    && curl -s -o /tmp/go.tar.gz https://storage.googleapis.com/golang/go$GOVERSION.linux-$ARCH.tar.gz \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \
    && rm /tmp/go.tar.gz

RUN mkdir -p /home/root/src \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | tee /etc/apt/sources.list.d/google-chrome.list \
    && apt-get update -qq \
    && apt-get install -y google-chrome-stable

RUN rm -rf /tmp/* /usr/local/bin/chromedriver

# 緊急対応
# パッケージがないため互換性のあるバージョンをハードコード
RUN curl -sS -o /tmp/chromedriver-linux64.zip https://edgedl.me.gvt1.com/edgedl/chrome/chrome-for-testing/116.0.5845.96/linux64/chromedriver-linux64.zip \
    && unzip /tmp/chromedriver-linux64.zip \
    && mv /chromedriver-linux64/chromedriver /usr/local/bin/

WORKDIR /go/src
