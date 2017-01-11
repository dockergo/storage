FROM alpine:3.4
MAINTAINER zhangdongdong1<zhangcong@kingsoft.com>

RUN mkdir -p /wpsep/storage/emptyfile \
&& apk update && apk add --no-cache nfs-utils

ADD ./need /wpsep/storage
COPY ./bin/storage /wpsep/storage
RUN chmod +x /wpsep/storage/start.sh /wpsep/storage/storage

EXPOSE 8080

ENTRYPOINT ["/wpsep/storage/start.sh"]
