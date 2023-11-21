package oauth2wrap

import "golang.org/x/oauth2"

type EndpointM struct {
	AuthURL       string `json:"auth_url,omitempty"`
	DeviceAuthURL string `json:"device_auth_url,omitempty"`
	TokenURL      string `json:"token_url,omitempty"`

	// AuthStyle optionally specifies how the endpoint wants the
	// client ID & client secret sent. The zero value means to
	// auto-detect.
	AuthStyle oauth2.AuthStyle `json:"auth_style,omitempty"`
}

type Conf struct {
	ClientID     string     `json:"id"`
	ClientSecret string     `json:"secret"`
	Endpoint     *EndpointM `json:"endpoint"`
	Scopes       []string   `json:"scopes,omitempty"`
	Api          string     `json:"api"`
}
