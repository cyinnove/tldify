package main

import (
	"fmt"
	"github.com/cyinnove/tldify"
)

func main() {
	url, err := tldify.Parse("http://sub.example.co.uk:8080/path?query=1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Subdomain:", url.Subdomain)
	fmt.Println("Domain:", url.Domain)
	fmt.Println("TLD:", url.TLD)
	fmt.Println("Port:", url.Port)
	fmt.Println("ICANN:", url.ICANN)
}