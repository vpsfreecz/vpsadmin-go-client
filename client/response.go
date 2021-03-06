package client

// Envelope represents a response from the API server
type Envelope struct {
	// Determines action success
	Status bool

	// Error message is Status is false
	Message string

	// Errors for individual parameters if Status is false
	Errors map[string][]string
}
