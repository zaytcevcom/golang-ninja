package keycloakclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// Audience can be either a string or an array of strings in Keycloak responses.
type Audience []string

func (a *Audience) UnmarshalJSON(data []byte) error {
	// Try array of strings
	var arr []string
	if err := json.Unmarshal(data, &arr); err == nil {
		*a = arr
		return nil
	}
	// Try single string
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*a = []string{single}
		return nil
	}
	return fmt.Errorf("invalid audience format: %s", string(data))
}

type IntrospectTokenResult struct {
	Exp    int      `json:"exp"`
	Iat    int      `json:"iat"`
	Aud    Audience `json:"aud"`
	Active bool     `json:"active"`
}

// IntrospectToken implements
// https://www.keycloak.org/docs/latest/authorization_services/index.html#obtaining-information-about-an-rpt
func (c *Client) IntrospectToken(ctx context.Context, token string) (IntrospectionResult, error) {
	var resp IntrospectionResult

	form := url.Values{}
	form.Set("client_id", c.clientID)
	form.Set("client_secret", c.clientSecret)
	form.Set("token", token)

	r := c.auth(ctx).
		SetBody(form.Encode()).
		SetResult(&resp)

	_, err := r.Post(fmt.Sprintf("/realms/%s/protocol/openid-connect/token/introspect", c.realm))
	if err != nil {
		return IntrospectionResult{}, fmt.Errorf("keycloak introspect: %w", err)
	}
	return resp, nil
}

func (c *Client) auth(ctx context.Context) *resty.Request {
	r := c.cli.R().SetContext(ctx)
	// Use HTTP Basic authentication for client auth.
	if c.clientID != "" || c.clientSecret != "" {
		r.SetBasicAuth(c.clientID, c.clientSecret)
	}
	// OAuth2 token/introspection endpoints expect URL-encoded forms.
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return r
}
