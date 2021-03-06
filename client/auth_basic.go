package client

import (
	"net/http"
)

type BasicAuth struct {
	Username string
	Password string
}

func (auth *BasicAuth) Authenticate(request *http.Request) {
	request.SetBasicAuth(auth.Username, auth.Password)
}

// SetBasicAuthentication enables HTTP basic auth on the client
func (client *Client) SetBasicAuthentication(username string, password string) {
	client.Authentication = &BasicAuth{
		Username: username,
		Password: password,
	}
}
