package config

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func GetElasticsearchConnection() {
	config := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTICSEARCH_URL")},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	es, err := elasticsearch.NewClient(config)

	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(es.Info())
	}
}
