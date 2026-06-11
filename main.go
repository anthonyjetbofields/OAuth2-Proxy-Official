package main

import (
	"fmt"
	"github.com/anthonyjetbofields/OAuth2-Proxy-Official/providers"
)

func main() {
	_ = providers.Config{ReverseProxy: true}
	fmt.Println("OAuth2-Proxy initialized with reverse proxy support.")
}