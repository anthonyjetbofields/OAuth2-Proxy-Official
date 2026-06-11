package providers

import (
	"fmt"
	"net/http"
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

// GetExternalHost returns the host considering X-Forwarded-Host
func GetExternalHost(r *http.Request, cfg Config) string {
	if cfg.ReverseProxy {
		if host := r.Header.Get("X-Forwarded-Host"); host != "" {
			return host
		}
	}
	return r.Host
}

// GetRedirectURI constructs the redirect URI
func GetRedirectURI(r *http.Request, cfg Config, path string) string {
	proto := GetExternalProto(r, cfg)
	host := GetExternalHost(r, cfg)
	return fmt.Sprintf("%s://%s%s", proto, host, path)
}

// ValidateState checks the state and logs mismatches
func ValidateState(expected, actual string, r *http.Request, cfg Config) bool {
	if expected != actual {
		fmt.Printf("DEBUG: State mismatch. Expected: %s, Received: %s, RedirectURI: %s\n", expected, actual, GetRedirectURI(r, cfg, r.URL.Path))
		return false
	}
	return true
}

// MakeStateCookie creates a state cookie, correctly setting the Secure flag
func MakeStateCookie(r *http.Request, cfg Config, name, value string) *http.Cookie {
	proto := GetExternalProto(r, cfg)
	secure := proto == "https"
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Secure:   secure,
		HttpOnly: true,
		Path:     "/",
	}
}
