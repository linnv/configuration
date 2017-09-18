FROM ubuntu:latest
MAINTAINER Jialin Wu <hi@jialinwu.com>

COPY 163.repos /etc/apt/sources.list
COPY ./base.sh /base.sh
COPY ./init.vim /init.vim

ENV	LC_ALL=C \
	LANGUAGE=C \
	LANG=C

RUN apt-get update && apt-get -y upgrade && apt-get -y install wget

RUN apt-get -y --no-install-recommends install \
  aptitude 

RUN  bash /base.sh
