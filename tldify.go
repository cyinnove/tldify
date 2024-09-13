
// Package tldify provides an extended version of net/url.URL.
// In addition to the standard URL parsing, tldify.URL includes additional fields 
// for Subdomain, Domain, TLD, Port, and ICANN status, offering more detailed URL analysis.
package tldify

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// URL is an extended version of net/url.URL. 
// It contains additional fields such as Subdomain, Domain, TLD, Port, and ICANN to store 
// more detailed information about the parsed URL.
type URL struct {
	Subdomain string // The subdomain part of the URL (e.g., "sub" in "sub.example.com").
	Domain    string // The registered domain name (e.g., "example" in "example.com").
	TLD       string // The top-level domain (TLD) (e.g., "com" in "example.com").
	Port      string // The port number if specified (e.g., "8080" in "example.com:8080").
	ICANN     bool   // Indicates if the TLD is recognized by ICANN.
	*url.URL         // Embeds net/url.URL to retain the standard URL structure.
}

// Parse is an extended version of net/url.Parse. 
// It parses the given URL string and returns a tldify.URL containing the standard URL fields
// along with additional details like Subdomain, Domain, TLD, Port, and ICANN status.
func Parse(s string) (*URL, error) {

	if !strings.HasPrefix(s, "http://") || !strings.HasPrefix(s, "https://"){
		s = fmt.Sprintf("http://%s", s)
	}

	parsedURL, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	// If no host is present, return the parsed URL as-is.
	if parsedURL.Host == "" {
		return &URL{URL: parsedURL}, nil
	}

	// Extract domain and port
	domain, port := splitDomainAndPort(parsedURL.Host)

	// Retrieve the effective top-level domain plus one (eTLD+1) and ICANN status
	etld1, err := publicsuffix.EffectiveTLDPlusOne(domain)
	suffix, icann := publicsuffix.PublicSuffix(strings.ToLower(domain))

	// Handle cases where the domain is valid but not registered with ICANN
	if err != nil && !icann && suffix == domain {
		etld1 = domain
		err = nil
	}
	if err != nil {
		return nil, err
	}

	// Split eTLD+1 into domain name and TLD
	i := strings.Index(etld1, ".")
	if i < 0 {
		return nil, fmt.Errorf("tldify: failed parsing %q", s)
	}
	domainName := etld1[:i]
	tld := etld1[i+1:]

	// Extract subdomain if present
	subdomain := ""
	if rest := strings.TrimSuffix(domain, "."+etld1); rest != domain {
		subdomain = rest
	}

	// Return the parsed URL with additional fields populated
	return &URL{
		Subdomain: subdomain,
		Domain:    domainName,
		TLD:       tld,
		Port:      port,
		ICANN:     icann,
		URL:       parsedURL,
	}, nil
}

// splitDomainAndPort splits the host into the domain and port components.
// If no port is present, an empty string is returned for the port.
func splitDomainAndPort(host string) (string, string) {
	for i := len(host) - 1; i >= 0; i-- {
		if host[i] == ':' {
			return host[:i], host[i+1:]
		} else if host[i] < '0' || host[i] > '9' {
			return host, ""
		}
	}
	// Fallback to return the host as-is if it's entirely numeric (which shouldn't happen)
	return host, ""
}
