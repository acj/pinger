# pinger

Pings a host reapeatedly and reports the round-trip times to Datadog.

## Getting Started

#### Bare Metal

```
$ go get github.com/DataDog/datadog-go/statsd
$ go get github.com/tatsushid/go-fastping
$ go build pinger.go
$ ./pinger
```

If you're running pinger as a non-root user, you may need to use `sudo ./pinger` because the ICMP pings require raw socket access.

#### Docker

```
$ docker pull acjensen/pinger
$ docker run -it pinger
```

### Configuring

```
$ ./pinger -h
Usage of ./pinger:
  -interval int
        Ping interval in milliseconds. Also serves as timeout. (default 2000)
  -metric string
        Metric name (default "udp_ping")
  -namespace string
        StatsD namespace for metrics (default "pinger.")
  -ping_host string
        The host to ping (default "localhost")
  -statsd_host string
        The StatsD host and port to use for metrics (default "0.0.0.0:8125")
```

Options can be passed to Docker with `-e`, for example:

```
$ docker run -it -e STATSD_HOST=192.168.1.10:8125 pinger
```
