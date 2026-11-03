package client_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

type oauthTransport func(*http.Request) (*http.Response, error)

func (transport oauthTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	return transport(request)
}

func TestOAuthRevocationRequiresTrustedOrigin(t *testing.T) {
	api := client.New("https://api.vpsfree.cz")
	api.SetExistingOAuth2Auth("local-test-token")
	dispatched := false
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(request *http.Request) (*http.Response, error) {
		dispatched = true
		return nil, nil
	})})

	err := api.RevokeAccessToken()
	if err == nil || !strings.Contains(err.Error(), "outside client origin") {
		t.Fatalf("expected untrusted-origin error, got %v", err)
	}
	if dispatched {
		t.Fatal("untrusted revocation reached the HTTP transport")
	}
	if api.Authentication == nil {
		t.Fatal("rejected revocation cleared authentication")
	}
}

func TestOAuthRevocationPreservesProductionEndpoint(t *testing.T) {
	for _, status := range []int{http.StatusOK, http.StatusServiceUnavailable} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			api := client.New("https://api.vpsfree.cz")
			api.SetExistingOAuth2Auth("local-test-token")
			if err := api.AllowOAuth2Origin("https://auth.vpsfree.cz"); err != nil {
				t.Fatal(err)
			}
			dispatched := false
			api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(request *http.Request) (*http.Response, error) {
				dispatched = true
				if request.Method != http.MethodPost || request.URL.String() != "https://auth.vpsfree.cz/_auth/oauth2/revoke" {
					t.Fatalf("unexpected revocation request: %s %s", request.Method, request.URL)
				}
				if err := request.ParseForm(); err != nil {
					t.Fatal(err)
				}
				if request.PostForm.Get("token") != "local-test-token" || request.Header.Get("X-HaveAPI-OAuth2-Token") != "local-test-token" {
					t.Fatal("revocation did not use the configured test token")
				}
				return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(""))}, nil
			})})

			err := api.RevokeAccessToken()
			if !dispatched {
				t.Fatal("trusted revocation did not reach the HTTP transport")
			}
			if status == http.StatusOK {
				if err != nil || api.Authentication != nil {
					t.Fatalf("successful revocation must clear authentication: %v", err)
				}
			} else if err == nil || api.Authentication == nil {
				t.Fatal("failed revocation must retain authentication")
			}
		})
	}
}
