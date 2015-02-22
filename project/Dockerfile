FROM google/debian:wheezy
MAINTAINER Kazuhiro Ogura <goura@goura.com>
RUN apt-get update && apt-get install -y ca-certificates
RUN mkdir /spotmc-ctrl
WORKDIR /spotmc-ctrl

ADD spotmc-ctrl /spotmc-ctrl/

EXPOSE 5000
CMD ["/spotmc-ctrl/spotmc-ctrl"]
