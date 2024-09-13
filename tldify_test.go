package tldify

import (
	"testing"

)

// Helper function to run the test cases
func test(input, sub, dom, tld string, icann, errorExpected bool, t *testing.T) {
	u, err := Parse(input)

	if err != nil && errorExpected {
		return
	} else if err != nil {
		t.Errorf("errored '%s'", err)
	} else if u.TLD != tld {
		t.Errorf("should have TLD '%s', got '%s'", tld, u.TLD)
	} else if u.Domain != dom {
		t.Errorf("should have Domain '%s', got '%s'", dom, u.Domain)
	} else if u.Subdomain != sub {
		t.Errorf("should have Subdomain '%s', got '%s'", sub, u.Subdomain)
	} else if u.ICANN != icann {
		t.Errorf("should have ICANN '%t', got '%t'", icann, u.ICANN)
	}
}

// Using t.Run for each test case
func TestParseURLs(t *testing.T) {
	t.Run("Basic domain without subdomain", func(t *testing.T) {
		test("http://foo.com", "", "foo", "com", true, false, t)
	})

	t.Run("Domain with multiple subdomains", func(t *testing.T) {
		test("http://zip.zop.foo.com", "zip.zop", "foo", "com", true, false, t)
	})

	t.Run("Domain with country-code TLD", func(t *testing.T) {
		test("http://au.com.au", "", "au", "com.au", true, false, t)
	})

	t.Run("Domain with complex subdomain and port", func(t *testing.T) {
		test("http://im.from.england.co.uk:1900", "im.from", "england", "co.uk", true, false, t)
	})

	t.Run("Simple HTTPS domain", func(t *testing.T) {
		test("https://google.com", "", "google", "com", true, false, t)
	})

	t.Run("Non-ICANN managed TLD", func(t *testing.T) {
		test("https://foo.notmanaged", "", "foo", "notmanaged", false, false, t)
	})

	t.Run("Capitalized TLD", func(t *testing.T) {
		test("https://google.Com", "", "google", "Com", true, false, t)
	})

	t.Run("Generic non-ICANN TLD", func(t *testing.T) {
		test("https://github.io", "", "github", "io", false, false, t)
	})

	// Test cases that expect errors
	t.Run("URL without dot", func(t *testing.T) {
		test("https://no_dot_should_not_panic", "", "", "", false, true, t)
	})

	t.Run("URL starts with dot", func(t *testing.T) {
		test("https://.start_with_dot_should_fail", "", "", "", false, true, t)
	})

	t.Run("URL ends with dot", func(t *testing.T) {
		test("https://ends_with_dot_should_fail.", "", "", "", false, true, t)
	})
}
