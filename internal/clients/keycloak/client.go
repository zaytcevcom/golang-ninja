package keycloakclient

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

//go:generate options-gen -out-filename=client_options.gen.go -from-struct=Options
type Options struct {
	basePath     string `option:"mandatory"`
	realm        string `option:"mandatory"`
	clientID     string `option:"mandatory"`
	clientSecret string `option:"mandatory"`
	debugMode    bool   `option:"optional"`
}

// Client is a tiny client to the KeyCloak realm operations. UMA configuration:
// http://localhost:3010/realms/Bank/.well-known/uma2-configuration
type Client struct {
	realm        string
	clientID     string
	clientSecret string
	cli          *resty.Client
}

// TokenResponse is a minimal subset of Keycloak token response used in tests.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

// IntrospectionResult reflects the RFC 7662 active flag.
type IntrospectionResult struct {
	Active bool `json:"active"`
}

func New(opts Options) (*Client, error) {
	if err := opts.Validate(); err != nil {
		return nil, fmt.Errorf("validate options: %v", err)
	}

	cli := resty.New()
	cli.SetDebug(opts.debugMode)
	cli.SetBaseURL(opts.basePath)

	return &Client{
		realm:        opts.realm,
		clientID:     opts.clientID,
		clientSecret: opts.clientSecret,
		cli:          cli,
	}, nil
}
