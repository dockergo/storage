FROM alpine:3.4
MAINTAINER zhangdongdong1<zhangcong@kingsoft.com>

RUN mkdir -p /wpsep/agent/emptyfile \
&& apk update && apk add --no-cache nfs-utils

ADD ./emptyfile /wpsep/agent/emptyfile
ADD ./need /wpsep/agent
COPY ./bin/agent /wpsep/agent
RUN chmod +x /wpsep/agent/start.sh /wpsep/agent/agent

EXPOSE 20808

ENTRYPOINT ["/wpsep/agent/start.sh"]
