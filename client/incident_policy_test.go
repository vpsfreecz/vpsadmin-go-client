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

func TestGenericIncidentDiscoveryAndExplicitFilters(t *testing.T) {
	api := client.New("https://api.example.test")
	var queries []url.Values
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		queries = append(queries, req.URL.Query())
		body := `{"status":true,"response":{"vpses":[{"id":1}],"vps":{"id":1}}}`
		if req.URL.Path == "/v7.0/incident_policies" {
			body = `{"status":true,"response":{"incident_policies":[{"key":"memory_pressure2","label":"Memory pressure"}]}}`
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})})
	registry, err := api.IncidentPolicy.Index.Prepare().Call()
	if err != nil {
		t.Fatal(err)
	}
	if registry.Output[0].Key != "memory_pressure2" {
		t.Fatalf("runtime key lost: %#v", registry.Output)
	}
	if _, err := api.Vps.Index.Prepare().Call(); err != nil {
		t.Fatal(err)
	}
	if _, err := api.Vps.Show.Prepare().SetPathParamInt("vps_id", 1).Call(); err != nil {
		t.Fatal(err)
	}
	for _, q := range queries {
		if _, ok := q["vps[incident_policy]"]; ok {
			t.Fatalf("injected selector: %v", q)
		}
	}
	selected := api.Vps.Index.Prepare()
	selected.NewInput().SetIncidentPolicy(registry.Output[0].Key).SetIncidentStatus("near_limit").SetIncidentEnforcementEnabled(false)
	if _, err := selected.Call(); err != nil {
		t.Fatal(err)
	}
	want := url.Values{"vps[incident_policy]": {"memory_pressure2"}, "vps[incident_status]": {"near_limit"}, "vps[incident_enforcement_enabled]": {"0"}}
	if !reflect.DeepEqual(queries[3], want) {
		t.Fatalf("selected query: %v", queries[3])
	}
}

func TestPolicyUpdateReturnsAuthoritativeRowAndPreservesNull(t *testing.T) {
	api := client.New("https://api.example.test")
	var bodies []map[string]interface{}
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/v7.0/vpses/1/incident_policies/memory_pressure2" || req.Method != "PUT" {
			t.Fatalf("policy route: %s %s", req.Method, req.URL)
		}
		var body map[string]interface{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		bodies = append(bodies, body)
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{"incident_policy":{"key":"memory_pressure2","configuration_valid":true,"warning_limit":5,"stop_limit":12,"stop_limit_override":12}}}`))}, nil
	})})
	update := api.Vps.IncidentPolicy.Update.Prepare().SetPathParamInt("vps_id", 1).SetPathParamString("key", "memory_pressure2")
	update.NewInput().SetWarningLimitOverrideNil(true).SetStopLimitOverride(12)
	response, err := update.Call()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]interface{}{"_meta": map[string]interface{}{}, "incident_policy": map[string]interface{}{"warning_limit_override": nil, "stop_limit_override": float64(12)}}
	if !reflect.DeepEqual(bodies[0], want) {
		t.Fatalf("policy body: %#v", bodies[0])
	}
	if response.Output.Key != "memory_pressure2" || !response.Output.ConfigurationValid || response.Output.StopLimit != 12 {
		t.Fatalf("authoritative response: %#v", response.Output)
	}
	reset := api.Vps.IncidentPolicy.Update.Prepare().SetPathParamInt("vps_id", 1).SetPathParamString("key", "memory_pressure2")
	reset.NewInput().SetStopLimitOverrideNil(true)
	if _, err := reset.Call(); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(bodies[1]["incident_policy"], map[string]interface{}{"stop_limit_override": nil}) {
		t.Fatalf("reset body: %#v", bodies[1])
	}
}

func TestSharedIncidentPermissionAndCountersStayOnVps(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		if req.Method == "PUT" {
			var body map[string]interface{}
			if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			want := map[string]interface{}{"_meta": map[string]interface{}{}, "vps": map[string]interface{}{"incident_enforcement_enabled": false}}
			if !reflect.DeepEqual(body, want) {
				t.Fatalf("shared permission body: %#v", body)
			}
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{}}`))}, nil
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{"vps":{"id":1,"incident_enforcement_enabled":false,"incident_warning_count":2,"incident_stop_threshold_count":1}}}`))}, nil
	})})
	update := api.Vps.Update.Prepare().SetPathParamInt("vps_id", 1)
	update.NewInput().SetIncidentEnforcementEnabled(false)
	ack, err := update.Call()
	if err != nil {
		t.Fatal(err)
	}
	if ack.Response.Meta != nil {
		t.Fatal("synchronous acknowledgement invented action state")
	}
	result, err := api.Vps.Show.Prepare().SetPathParamInt("vps_id", 1).Call()
	if err != nil {
		t.Fatal(err)
	}
	if result.Output.IncidentEnforcementEnabled || result.Output.IncidentWarningCount != 2 || result.Output.IncidentStopThresholdCount != 1 {
		t.Fatalf("shared readback: %#v", result.Output)
	}
}

func TestInvalidCurrentPolicyScalarNullsAndSavedHistory(t *testing.T) {
	var current client.ActionVpsIncidentPolicyShowOutput
	if err := json.Unmarshal([]byte(`{"key":"memory_pressure2","configuration_valid":false,"warning_limit":null,"stop_limit":null,"status":null,"count_7d":2,"count_30d":3,"evaluation_report_id":12,"evaluated_at":"2026-11-18T12:00:00Z","evaluation_state":"recorded"}`), &current); err != nil {
		t.Fatal(err)
	}
	// The retained generator decodes missing/null scalar fields to their zero value.
	if current.ConfigurationValid || current.WarningLimit != 0 || current.StopLimit != 0 || current.Status != "" || current.Count7d != 2 || current.Count30d != 3 {
		t.Fatalf("current scalar facts: %#v", current)
	}
	if current.EvaluationReportId != 12 || current.EvaluatedAt != "2026-11-18T12:00:00Z" || current.EvaluationState != "recorded" {
		t.Fatalf("historical pointer: %#v", current)
	}
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/v7.0/incident_reports/12/policy_evaluations" {
			t.Fatalf("history route: %s", req.URL)
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{"policy_evaluations":[{"policy_key":"memory_pressure2","state":"recorded","count_7d":5,"warning_limit":5,"stop_limit":10}]}}`))}, nil
	})})
	history, err := api.IncidentReport.PolicyEvaluation.Index.Prepare().SetPathParamInt("incident_report_id", 12).Call()
	if err != nil {
		t.Fatal(err)
	}
	if history.Output[0].PolicyKey != "memory_pressure2" || history.Output[0].WarningLimit != 5 || history.Output[0].StopLimit != 10 {
		t.Fatalf("saved evaluation: %#v", history.Output)
	}
}

func TestIncidentNoticeMetadataThroughGeneratedActions(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		body := `{"status":true,"response":{"incident_reports":[{"id":1,"notice_kind":null,"trigger_incident_report_id":null,"notice_mail_state":null,"notice_policy_key":null,"notice_policy_label":null}]}}`
		if strings.HasSuffix(req.URL.Path, "/2") {
			body = `{"status":true,"response":{"incident_report":{"id":2,"notice_kind":"stop_completed","trigger_incident_report_id":100,"notice_mail_state":"queued","notice_policy_key":"retired_policy","notice_policy_label":"Saved label"}}}`
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})})
	page, err := api.IncidentReport.Index.Prepare().Call()
	if err != nil {
		t.Fatal(err)
	}
	ordinary := page.Output[0]
	if ordinary.NoticeKind != "" || ordinary.TriggerIncidentReportId != 0 || ordinary.NoticeMailState != "" {
		t.Fatalf("ordinary null scalar values: %#v", ordinary)
	}
	result, err := api.IncidentReport.Show.Prepare().SetPathParamInt("incident_report_id", 2).Call()
	if err != nil {
		t.Fatal(err)
	}
	if result.Output.NoticeKind != "stop_completed" || result.Output.TriggerIncidentReportId != 100 || result.Output.NoticeMailState != "queued" || result.Output.NoticePolicyKey != "retired_policy" || result.Output.NoticePolicyLabel != "Saved label" {
		t.Fatalf("notice metadata lost: %#v", result.Output)
	}
	for _, state := range []string{"pending", "queued", "sent", "failed", "uncertain", "suppressed"} {
		var out client.ActionIncidentReportShowOutput
		if err := json.Unmarshal([]byte(`{"notice_kind":"warning","trigger_incident_report_id":null,"notice_mail_state":"`+state+`"}`), &out); err != nil {
			t.Fatal(err)
		}
		if out.NoticeKind != "warning" || out.TriggerIncidentReportId != 0 || out.NoticeMailState != state {
			t.Fatalf("mail state/nullable link: %#v", out)
		}
	}
}

func TestSavedMembershipScalarNullsAndPeriods(t *testing.T) {
	for _, tc := range []struct {
		wire    string
		matched bool
	}{
		{`{"matched":true,"warning_period_days":7,"stop_period_days":30}`, true},
		{`{"matched":false,"warning_period_days":7,"stop_period_days":30}`, false},
		{`{"matched":null,"warning_period_days":7,"stop_period_days":30}`, false},
	} {
		var out client.ActionIncidentReportPolicyEvaluationIndexOutput
		if err := json.Unmarshal([]byte(tc.wire), &out); err != nil {
			t.Fatal(err)
		}
		if out.Matched != tc.matched || out.WarningPeriodDays != 7 || out.StopPeriodDays != 30 {
			t.Fatalf("saved membership or periods: %#v", out)
		}
	}
	var missing client.ActionIncidentReportPolicyEvaluationIndexOutput
	if err := json.Unmarshal([]byte(`{}`), &missing); err != nil {
		t.Fatal(err)
	}
	if missing.Matched || missing.WarningPeriodDays != 0 || missing.StopPeriodDays != 0 {
		t.Fatalf("missing scalar behavior changed: %#v", missing)
	}
}
