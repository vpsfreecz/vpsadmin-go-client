package client

import (
	"errors"
	"net/http"
)

type Mode int

const (
	// Send the token via HTTP header X-HaveAPI-Auth-Token
	HttpHeader Mode = iota

	// Send the token via query parameter _auth_token
	QueryParameter = iota
)

type TokenAuth struct {
	// Resource for token manipulation
	Resource *AuthTokenResourceToken

	// The authentication token
	Token string

	// Mode determines how is the authentication token sent to the API
	Mode Mode
}

func (auth *TokenAuth) Authenticate(request *http.Request) {
	switch auth.Mode {
	case HttpHeader:
		request.Header.Set("X-HaveAPI-Auth-Token", auth.Token)

	case QueryParameter:
		q := request.URL.Query()
		q.Add("_auth_token", auth.Token)
		request.URL.RawQuery = q.Encode()
	}
}

// SetNewTokenAuth obtains a new authentication token with username and password
//
// lifetime can be one of: fixed, renewable_auto, renewable_manual or permanent
// interval determines how many seconds will the token be valid
func (client *Client) SetNewTokenAuth(username string, password string, lifetime string, interval int) error {
	resource := NewAuthTokenResourceToken(client)

	request := resource.Request.Prepare()
	request.SetInput(&AuthTokenActionTokenRequestInput{
		Login: username,
		Password: password,
		Lifetime: lifetime,
		Interval: int64(interval),
	})

	resp, err := request.Call()

	if err != nil {
		return err
	} else if !resp.Status {
		return errors.New("Unable to request token: " + resp.Message)
	}

	client.Authentication = &TokenAuth{
		Resource: resource,
		Token: resp.Output.Token,
		Mode: HttpHeader,
	}

	return nil
}

// SetExistingTokenAuth will use a previously acquired token
func (client *Client) SetExistingTokenAuth(token string) {
	client.Authentication = &TokenAuth{
		Resource: NewAuthTokenResourceToken(client),
		Token: token,
		Mode: HttpHeader,
	}
}

// SetTokenAuthMode can be used to change the way the token is sent to the API
func (client *Client) SetTokenAuthMode(mode Mode) {
	client.Authentication.(*TokenAuth).Mode = mode
}

// RevokeAuthToken will revoke the authentication token and remove authentication
// from the client
func (client *Client) RevokeAuthToken() error {
	revoke := client.Authentication.(*TokenAuth).Resource.Revoke.Prepare()
	resp, err := revoke.Call()

	if err != nil {
		return err
	} else if !resp.Status {
		return errors.New("Unable to revoke token: " + resp.Message)
	}

	client.Authentication = nil
	return nil
}
