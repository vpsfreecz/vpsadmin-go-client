package client_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func TestSystemConfigNullableOutput(t *testing.T) {
	for _, tc := range []struct {
		name    string
		payload string
		want    bool
	}{
		{"true", `{"choices":[{"value":false,"label":"No"}],"value_nullable":true}`, true},
		{"false", `{"choices":[{"value":0,"label":"Zero"}],"value_nullable":false}`, false},
		{"null", `{"choices":null,"value_nullable":null}`, false},
		{"absent", `{}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			api := client.New("https://api.example.test")
			calls := 0
			api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != "GET" || req.URL.Path != "/v7.0/system_configs/core/mode" {
					t.Fatalf("route: %s %s", req.Method, req.URL)
				}
				body := `{"status":true,"response":{"system_config":` + tc.payload + `}}`
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
			})})
			result, err := api.SystemConfig.Show.Prepare().
				SetPathParamString("category", "core").SetPathParamString("name", "mode").Call()
			if err != nil {
				t.Fatal(err)
			}
			if !result.Status || result.Output == nil || result.Output.ValueNullable != tc.want || calls != 1 {
				t.Fatalf("nullable output: %#v; calls=%d", result, calls)
			}
		})
	}
}
