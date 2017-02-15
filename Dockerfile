FROM alpine:3.5
MAINTAINER flyaways<meagleglass@gmail.com>

RUN mkdir -p /go/resources

ADD ./go/* ./bin/agent /go

ADD ./go/resources /go/resources

RUN chmod +x /go/start.sh /go/agent

EXPOSE 8080

ENTRYPOINT ["/go/start.sh"]
