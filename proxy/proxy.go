//
// Package proxy implements a HTTP caching proxy for Node package registry (NPM).
// See https://github.com/pkgems/npm-cache-proxy/ for more information about proxy.
//
package proxy

import (
	"net/http"
	"time"
)

// Proxy is the proxy instance, it contains Database and HttpClient as static options
// and GetOptions as dynamic options provider
type Proxy struct {
	Database   Database
	HttpClient *http.Client
}

// Options provides dynamic options for Proxy.
// This can be used for namespace separation,
// allowing multiple users use the same proxy instance simultaneously.
type Options struct {
	DatabasePrefix     string
	DatabaseExpiration time.Duration
	UpstreamAddress    string
	AuthToken          string
}

// Database provides interface for data storage.
type Database interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
	Delete(key string) error
	Keys(prefix string) ([]string, error)
	Health() error
}
