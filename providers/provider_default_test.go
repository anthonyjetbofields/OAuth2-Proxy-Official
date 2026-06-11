package providers

import (
	"net/http/httptest"
	"testing"
)

func TestGetRedirectURI(t *testing.T) {
	req := httptest.NewRequest("GET", "http://internal:8080/oauth2/callback", nil)
	req.Header.Set("X-Forwarded-Proto", "https")
	req.Header.Set("X-Forwarded-Host", "external.com")

	cfg := Config{ReverseProxy: true}
	uri := GetRedirectURI(req, cfg, "/oauth2/callback")
	expected := "https://external.com/oauth2/callback"

	if uri != expected {
		t.Errorf("Expected redirect URI %s, got %s", expected, uri)
	}
}

func TestValidateState(t *testing.T) {
	req := httptest.NewRequest("GET", "http://internal:8080/oauth2/callback", nil)
	cfg := Config{ReverseProxy: true}
	
	if !ValidateState("state123", "state123", req, cfg) {
		t.Errorf("Expected state to be valid")
	}

	if ValidateState("state123", "wrongstate", req, cfg) {
		t.Errorf("Expected state to be invalid")
	}
}

func TestMakeStateCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "http://internal:8080/login", nil)
	req.Header.Set("X-Forwarded-Proto", "https")
	cfg := Config{ReverseProxy: true}

	cookie := MakeStateCookie(req, cfg, "_oauth2_proxy_state", "state123")
	
	if !cookie.Secure {
		t.Errorf("Expected cookie to be secure when X-Forwarded-Proto is https")
	}
	
	if !cookie.HttpOnly {
		t.Errorf("Expected cookie to be HttpOnly")
	}
}
