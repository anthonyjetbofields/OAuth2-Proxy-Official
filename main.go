package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Config holds proxy settings
type Config struct {
	ReverseProxy bool
}

// GetExternalProto returns the protocol considering X-Forwarded-Proto
func GetExternalProto(r *http.Request, cfg Config) string {
	if cfg.ReverseProxy {
		if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
			return proto
		}
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

// ValidateState checks the state and logs mismatches
func ValidateState(expected, actual string) bool {
	if expected != actual {
		fmt.Printf("DEBUG: State mismatch. Expected: %s, Received: %s\n", expected, actual)
		return false
	}
	return true
}

func main() {
	fmt.Println("OAuth2-Proxy initialized with reverse proxy support.")
}