package main

import (
	"net/http"
	"time"

	npmproxy "github.com/pkgems/npm-cache-proxy/proxy"
	"github.com/go-redis/redis"
)

func main() {
	proxy := npmproxy.Proxy{
		Database: npmproxy.DatabaseRedis{
			Client: redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				DB:       0,
				Password: "",
			}),
		},
		HttpClient: &http.Client{},
	}

	proxy.Server(npmproxy.ServerOptions{
		ListenAddress: "localhost:8080",
		GetOptions: func() (npmproxy.Options, error) {
			return npmproxy.Options{
				DatabasePrefix:     "ncp-",
				DatabaseExpiration: 1 * time.Hour,
				UpstreamAddress:    "https://registry.npmjs.org",
				AuthToken:          "",
			}, nil
		},
	}).ListenAndServe()
}
