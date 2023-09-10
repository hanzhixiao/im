package main

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "139639hh",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
		return
	}
	resp, err := client.Info()
	if err != nil {
		panic(err)
		return
	}
	log.Println(resp)
}
