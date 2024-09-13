# tldify

`tldify` is a Go package that extends the functionality of the standard `net/url` package by adding support for parsing URLs into additional fields such as Subdomain, Domain, TLD (Top-Level Domain), and Port. It is designed to simplify URL parsing when you need finer granularity beyond the default URL structure.

This package is inspired by [go-tld](https://github.com/jpillora/go-tld) by @jpillora.

## Features

- Parse URLs into components like subdomain, domain, TLD, and port.
- Support for ICANN-managed TLDs using `golang.org/x/net/publicsuffix`.
- Similar API to `net/url` for easy integration into existing projects.

## Installation

To install `tldify`, use `go get`:

```bash
go get github.com/cyinnove/tldify
```

## Usage

Here's an example of how to use `tldify`:

```go
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
```

Output:

```
Subdomain: sub
Domain: example
TLD: co.uk
Port: 8080
ICANN: true
```

`tldify` exposes a single main function `Parse`, which mirrors `net/url.Parse` but returns an enriched `URL` struct:

```go
func Parse(s string) (*URL, error)
```

### URL Struct

The `URL` struct embeds the standard `net/url.URL` and adds the following fields:

- `Subdomain`: The subdomain portion of the URL.
- `Domain`: The domain portion of the URL.
- `TLD`: The Top-Level Domain (TLD) portion of the URL.
- `Port`: The port specified in the URL (if any).
- `ICANN`: A boolean indicating whether the TLD is an ICANN-managed public suffix.



## Inspiration

This package was inspired by [go-tld](https://github.com/jpillora/go-tld), a powerful tool for extracting domains and TLDs from URLs.

## License

`tldify` is licensed under the MIT License. See the `LICENSE` file for more information.
