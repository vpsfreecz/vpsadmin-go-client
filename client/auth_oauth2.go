package client

import (
	"fmt"
	"net/http"
)

type OAuth2Auth struct {
	// The authentication token
	AccessToken string
}

func (auth *OAuth2Auth) Authenticate(request *http.Request) {
	request.Header.Set("X-HaveAPI-OAuth2-Token", auth.AccessToken)
}

// SetExistingTokenAuth will use a previously acquired access token
func (client *Client) SetExistingOAuth2Auth(accessToken string) {
	client.Authentication = &OAuth2Auth{
		AccessToken: accessToken,
	}
}

// RevokeAuthToken will revoke the access token and remove authentication
// from the client
func (client *Client) RevokeAccessToken() error {
	httpClient := &http.Client{}

	req, err := http.NewRequest("POST", "https://auth.vpsfree.cz/_auth/oauth2/revoke", nil)

	if err != nil {
		return err
	}

	if client.Authentication != nil {
		client.Authentication.Authenticate(req)
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Unable to revoke access token, HTTP %v", resp.StatusCode)
	}

	client.Authentication = nil
	return nil
}
