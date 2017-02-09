FROM alpine:3.5
MAINTAINER flyaways<meagleglass@gmail.com>

ADD ./go /go
COPY ./bin/agent /go
RUN chmod +x /go/start.sh /go/agent

EXPOSE 8080

ENTRYPOINT ["/go/start.sh"]
