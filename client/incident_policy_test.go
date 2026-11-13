package client_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func TestIncidentReadSelectorsAreOnlySentWhenSelected(t *testing.T) {
	api := client.New("https://api.example.test")
	var queries []url.Values
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		queries = append(queries, req.URL.Query())
		body := `{"status":true,"response":{"vpses":[{"id":1}],"vps":{"id":1}}}`
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})})
	if _, err := api.Vps.Index.Prepare().Call(); err != nil {
		t.Fatal(err)
	}
	if _, err := api.Vps.Show.Prepare().SetPathParamInt("vps_id", 1).Call(); err != nil {
		t.Fatal(err)
	}
	page := api.Vps.Index.Prepare()
	page.NewInput().SetLimit(20)
	if _, err := page.Call(); err != nil {
		t.Fatal(err)
	}
	for _, query := range queries {
		if _, present := query["vps[incident_policy]"]; present {
			t.Fatalf("initial read injected selector: %v", query)
		}
	}
	selected := api.Vps.Index.Prepare()
	selected.NewInput().SetIncidentPolicy("oomd").SetIncidentStatus("near_limit").SetIncidentEnforcementEnabled(false)
	if _, err := selected.Call(); err != nil {
		t.Fatal(err)
	}
	want := url.Values{"vps[incident_policy]": {"oomd"}, "vps[incident_status]": {"near_limit"}, "vps[incident_enforcement_enabled]": {"0"}}
	if !reflect.DeepEqual(queries[3], want) {
		t.Fatalf("selected query: %v", queries[3])
	}
}

func TestIncidentSettingsPreserveNullFalseAndIsolatedFields(t *testing.T) {
	api := client.New("https://api.example.test")
	var bodies []map[string]interface{}
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		var body map[string]interface{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		bodies = append(bodies, body)
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{}}`))}, nil
	})})
	update := api.Vps.Update.Prepare().SetPathParamInt("vps_id", 1)
	update.NewInput().SetIncidentPolicy("oomd").SetIncidentWarningLimitOverrideNil(true).SetIncidentStopLimitOverride(12).SetIncidentEnforcementEnabled(false)
	response, err := update.Call()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]interface{}{"_meta": map[string]interface{}{}, "vps": map[string]interface{}{"incident_policy": "oomd", "incident_warning_limit_override": nil, "incident_stop_limit_override": float64(12), "incident_enforcement_enabled": false}}
	if !reflect.DeepEqual(bodies[0], want) {
		t.Fatalf("settings body: %#v", bodies[0])
	}
	if response.Response.Meta != nil {
		t.Fatal("synchronous update invented action metadata")
	}
	reset := api.Vps.Update.Prepare().SetPathParamInt("vps_id", 1)
	reset.NewInput().SetIncidentPolicy("oomd").SetIncidentStopLimitOverrideNil(true)
	if _, err := reset.Call(); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(bodies[1], map[string]interface{}{"_meta": map[string]interface{}{}, "vps": map[string]interface{}{"incident_policy": "oomd", "incident_stop_limit_override": nil}}) {
		t.Fatalf("reset body: %#v", bodies[1])
	}
	plain := api.Vps.Update.Prepare().SetPathParamInt("vps_id", 1)
	plain.NewInput().SetHostname("ordinary-update")
	if _, err := plain.Call(); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(bodies[2], map[string]interface{}{"_meta": map[string]interface{}{}, "vps": map[string]interface{}{"hostname": "ordinary-update"}}) {
		t.Fatalf("unrelated update: %#v", bodies[2])
	}
}

func TestIncidentSummaryTypesAndServerRefusal(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		body := `{"status":true,"response":{"vps":{"incident_policy":"oomd","incident_policy_label":"VPS memory-pressure incidents","incident_count_7d":5,"incident_count_30d":10,"incident_warning_limit":5,"incident_stop_limit":10,"incident_warning_limit_override":null,"incident_stop_limit_override":12,"incident_enforcement_enabled":false,"incident_mode":"enforce","incident_stop_enforced":false,"incident_status":"stop_threshold","incident_last_incident_at":"2026-11-14T10:00:00Z","incident_last_warning_at":null,"incident_last_stop_at":null,"incident_next_stop_allowed_at":"2026-11-15T10:00:00Z","incident_accounting_since":"2026-11-01T00:00:00Z","incident_action_state":"pending","incident_action_chain_id":42}}}`
		if req.URL.Query().Get("vps[incident_policy]") == "unregistered" {
			body = `{"status":false,"message":"Invalid incident policy","errors":{"incident_policy":["is not included"]}}`
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})})
	show := api.Vps.Show.Prepare().SetPathParamInt("vps_id", 1)
	show.NewInput().SetIncidentPolicy("oomd")
	response, err := show.Call()
	if err != nil {
		t.Fatal(err)
	}
	want := &client.ActionVpsShowOutput{IncidentPolicy: "oomd", IncidentPolicyLabel: "VPS memory-pressure incidents", IncidentCount7d: 5, IncidentCount30d: 10, IncidentWarningLimit: 5, IncidentStopLimit: 10, IncidentStopLimitOverride: 12, IncidentMode: "enforce", IncidentStatus: "stop_threshold", IncidentLastIncidentAt: "2026-11-14T10:00:00Z", IncidentNextStopAllowedAt: "2026-11-15T10:00:00Z", IncidentAccountingSince: "2026-11-01T00:00:00Z", IncidentActionState: "pending", IncidentActionChainId: 42}
	if !reflect.DeepEqual(response.Output, want) {
		t.Fatalf("summary: %#v", response.Output)
	}
	// The retained Go generator represents null/missing scalar output as zero values.
	show.NewInput().SetIncidentPolicy("unregistered")
	refused, err := show.Call()
	if err != nil {
		t.Fatal(err)
	}
	if refused.Status || refused.Message != "Invalid incident policy" || refused.Output != nil {
		t.Fatalf("server refusal lost: %#v", refused)
	}
}
