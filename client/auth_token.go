package client

import (
	"fmt"
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

type TokenAuthOptions struct {
	Interval int64
	Lifetime string
	Password string
	User string
	TotpCallback func(input *AuthTokenActionTokenTotpInput) error
}

// SetNewTokenAuth obtains a new authentication token with login credentials
func (client *Client) SetNewTokenAuth(options *TokenAuthOptions) error {
	resource := NewAuthTokenResourceToken(client)
	auth := &TokenAuth{Resource: resource, Mode: HttpHeader}
	auth.setDefaultOptions(options)

	request := resource.Request.Prepare()
	request.SetInput(&AuthTokenActionTokenRequestInput{
		Interval: options.Interval,
		Lifetime: options.Lifetime,
		Password: options.Password,
		User: options.User,
	})

	resp, err := request.Call()

	if err != nil {
		return err
	} else if !resp.Status {
		return fmt.Errorf("Unable to request token: %v", resp.Message)
	}

	if resp.Output.Complete {
		auth.Token = resp.Output.Token
		client.Authentication = auth
		return nil
	}

	return auth.nextAuthenticationStep(
		options,
		resp.Output.NextAction,
		resp.Output.Token,
	);
}

func (auth *TokenAuth) setDefaultOptions(options *TokenAuthOptions) {
	if options.Lifetime == "" {
		options.Lifetime = "renewable_auto"
	}

	if options.Interval == 0 {
		options.Interval = 300
	}
}

// nextAuthenticationStep performs authentication steps recursively, until
// the authentication is completed
func (auth *TokenAuth) nextAuthenticationStep(options *TokenAuthOptions, action string, token string) error {
	if action == "totp" {
		action := auth.Resource.Totp.Prepare()
		input := action.NewInput()
		input.SetToken(token)

		if options.TotpCallback == nil {
			return fmt.Errorf("Implement callback TotpCallback")
		}

		if err := options.TotpCallback(input); err != nil {
			return fmt.Errorf("TotpCallback failed: %v", err)
		}

		resp, err := action.Call()

		if err != nil {
			return err
		} else if !resp.Status {
			return fmt.Errorf("Failed at authentication step '%s': %v", action, resp.Message)
		}

		if resp.Output.Complete {
			auth.Token = resp.Output.Token
			auth.Resource.Client.Authentication = auth
			return nil
		}

		return auth.nextAuthenticationStep(
			options,
			resp.Output.NextAction,
			resp.Output.Token,
		);
	}

	return fmt.Errorf("Unsupported authentication action '%s'", action)
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
		return fmt.Errorf("Unable to revoke token: %v", resp.Message)
	}

	client.Authentication = nil
	return nil
}
