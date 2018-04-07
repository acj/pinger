FROM golang:1.10.1-alpine

ARG ping_host=localhost
ARG statsd_host=0.0.0.0:8125
ARG interval=2000
ARG namespace=pinger.
ARG metric=udp_ping

ENV PING_HOST $ping_host
ENV STATSD_HOST $statsd_host
ENV INTERVAL $interval
ENV NAMESPACE $namespace
ENV METRIC $metric

RUN apk update && apk add git

RUN go get github.com/DataDog/datadog-go/statsd && \
    go get github.com/tatsushid/go-fastping

ADD pinger.go pinger.go

RUN go build pinger.go && \
    rm *.go

ENTRYPOINT ./pinger -ping_host $PING_HOST -statsd_host $STATSD_HOST -interval $INTERVAL -namespace $NAMESPACE -metric $METRIC
