package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token     string `envconfig:"SPACETRADERS_TOKEN"`
	Debug     bool   `envconfig:"DEBUG"`
	Port      string `envconfig:"PORT" default:"8080"`
	TargetURL string `envconfig:"TARGET_URL" default:"https://api.spacetraders.io"`
}

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Error loading configuration: %v\n", err)
	}
	log.Print(config)
	proxy(config)
}

func proxyHandler(config Config) http.Handler {
	target, err := url.Parse(config.TargetURL)
	if err != nil {
		log.Fatal("Error parsing target URL:", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	throttle := time.Tick(600 * time.Millisecond)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-throttle // Wait for the next tick to proceed

		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println("DUMP REQUEST: ", err)
		}
		if config.Debug {
			fmt.Println(string(requestDump))
		}
		if config.Token != "" {
			r.Header.Set("Authorization", "Bearer "+config.Token)
		}
		// log.Print("request from " + r.Header.Get("Ship-Symbol"))
		r.Host = target.Host
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		proxy.ServeHTTP(w, r)
	})
}

func proxy(config Config) {
	// Create a router and register the proxy handler
	router := http.NewServeMux()
	router.Handle("/", proxyHandler(config))

	// Start the server
	log.Println("Proxy server listening on " + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
