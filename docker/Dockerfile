FROM alpine:3.4
MAINTAINER flyaways<meagleglass@gmail.com>

ADD ./need /go
COPY ./bin/storage /go
RUN chmod +x /go/start.sh /wpsep/storage/storage

EXPOSE 8080

ENTRYPOINT ["/wpsep/storage/start.sh"]
