FROM alpine:3.5
MAINTAINER flyaways<meagleglass@gmail.com>

RUN mkdir -p /flyaways/resources /flyaways/data

ADD ./go/* ./bin/agent /flyaways

ADD ./go/resources/* /flyaways/resources

RUN chmod +x /flyaways/start.sh /flyaways/agent

EXPOSE 8080

WORKDIR /flyaways
ENTRYPOINT ["/flyaways/entrypoint.sh"]
