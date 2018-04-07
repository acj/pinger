package main

import (
	"flag"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"time"
)

func main() {
	var statsd_host string
	var ping_host string
	var interval int
	var namespace string
	var metric string

	flag.StringVar(&statsd_host, "statsd_host", "0.0.0.0:8125", "The StatsD host and port to use for metrics")
  flag.StringVar(&ping_host, "ping_host", "localhost", "The host to ping")
	flag.IntVar(&interval, "interval", 2000, "Ping interval in milliseconds. Also serves as timeout.")
	flag.StringVar(&namespace, "namespace", "pinger.", "StatsD namespace for metrics")
	flag.StringVar(&namespace, "metric", "udp_ping", "Metric name")
	flag.Parse()

	c, err := statsd.New(statsd_host)
	if err != nil {
		log.Fatal("Failed to create statsd instance: ", err.Error())
	}
	c.Namespace = namespace

	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ping_host)
	if err != nil {
		log.Fatal("Failed to resolve hostname: ", err.Error())
	}
	p.Size = 512
	p.MaxRTT = time.Duration(interval) * time.Millisecond
//	p.Network("udp")
	p.AddIPAddr(ra)

	for {
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
//			log.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
			err = c.Histogram(metric, rtt.Seconds() * 1000, nil, 1)
			if err != nil {
				log.Printf("Failed to report stat: ", err.Error())
			}
		}
		p.OnIdle = func() {
			log.Println("Tick")
		}
		err = p.Run()
		if err != nil {
			log.Println("Error executing ping: ", err.Error())
		}
	}
}
