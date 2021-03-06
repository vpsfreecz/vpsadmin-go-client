package client

import (
	"net/http"
)

// Authenticator is used to provide authentication for HTTP requests
type Authenticator interface {
	Authenticate(request *http.Request)
}
