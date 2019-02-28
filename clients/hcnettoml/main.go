package hcnettoml

import "net/http"

// HcnetTomlMaxSize is the maximum size of hcnet.toml file
const HcnetTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the hcnet.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/hcnet.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a Hcnet.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a Hcnet.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

type ClientInterface interface {
	GetHcnetToml(domain string) (*Response, error)
	GetHcnetTomlByAddress(addy string) (*Response, error)
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a hcnet.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetHcnetToml returns hcnet.toml file for a given domain
func GetHcnetToml(domain string) (*Response, error) {
	return DefaultClient.GetHcnetToml(domain)
}

// GetHcnetTomlByAddress returns hcnet.toml file of a domain fetched from a
// given address
func GetHcnetTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetHcnetTomlByAddress(addy)
}
