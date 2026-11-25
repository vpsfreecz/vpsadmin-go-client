package client_test

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func TestDisableBackupsDefaultsOptOutAndOrdinaryReceipt(t *testing.T) {
	for _, tc := range []struct {
		name      string
		configure func(*client.ActionDatasetDisableBackupsInvocation)
		params    map[string]interface{}
	}{
		{"omitted", func(*client.ActionDatasetDisableBackupsInvocation) {}, map[string]interface{}{}},
		{"recursive-default-preservation", func(a *client.ActionDatasetDisableBackupsInvocation) { a.NewInput().SetRecursive(true) }, map[string]interface{}{"recursive": true}},
		{"explicit-opt-out", func(a *client.ActionDatasetDisableBackupsInvocation) {
			a.NewInput().SetBackupFirst(false).SetRecursive(false)
		}, map[string]interface{}{"backup_first": false, "recursive": false}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			api := client.New("https://api.example.test")
			calls := 0
			api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != "POST" || req.URL.Path != "/v7.0/datasets/7/disable_backups" {
					t.Fatalf("route: %s %s", req.Method, req.URL)
				}
				var body map[string]interface{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(body["dataset"], tc.params) {
					t.Fatalf("destructive input: %#v; want %#v", body, tc.params)
				}
				output := map[string]interface{}{"dataset": map[string]interface{}{}, "_meta": map[string]interface{}{"action_state_id": int64(41)}}
				payload, err := json.Marshal(map[string]interface{}{"status": true, "response": output})
				if err != nil {
					t.Fatal(err)
				}
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(string(payload)))}, nil
			})})
			action := api.Dataset.DisableBackups.Prepare().SetPathParamInt("dataset_id", 7)
			tc.configure(action)
			result, err := action.Call()
			if err != nil {
				t.Fatal(err)
			}
			if !result.Status {
				t.Fatalf("ordinary acknowledgement: %#v", result)
			}
			if result.Response.Meta == nil || result.Response.Meta.ActionStateId != 41 {
				t.Fatalf("ordinary receipt: %#v", result.Response.Meta)
			}
			if calls != 1 {
				t.Fatalf("one explicit mutation must submit once: %d calls", calls)
			}
		})
	}
}
