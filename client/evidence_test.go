package client_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

// Caller-owned rows are necessary: generated Go scalars collapse null to zero.
// The runtime accepts this complete envelope without modifying generated code.
type evidenceRow struct {
	CreditUs           *int64  `json:"credit_us"`
	FromAt             *string `json:"from_at"`
	ToAt               *string `json:"to_at"`
	Measurement        *string `json:"measurement"`
	PeakUnit           *string `json:"peak_unit"`
	PeakValue          *int64  `json:"peak_value"`
	PeakIsLowerBound   *bool   `json:"peak_is_lower_bound"`
	ConditionValue     *int64  `json:"condition_value"`
	ConfigurationValid bool    `json:"configuration_valid"`
	Vps                *struct {
		Id int64 `json:"id"`
	} `json:"vps"`
}
type evidenceMeta struct {
	Available            *bool    `json:"available"`
	ConfigurationValid   bool     `json:"configuration_valid"`
	UnavailableSignals   []string `json:"unavailable_signals"`
	TotalCount           int64    `json:"total_count"`
	NextOffset           *int64   `json:"next_offset"`
	ConditionMeasurement *string  `json:"condition_measurement"`
	ConditionOperator    *string  `json:"condition_operator"`
	ConditionValue       *int64   `json:"condition_value"`
	ConditionUnit        *string  `json:"condition_unit"`
	WindowComplete       *bool    `json:"window_complete"`
}

func TestEvidenceCustomResponsesPreserveNullZeroAndArrays(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetExistingTokenAuth("fixture-token")
	cases := []struct {
		body    string
		credit  *int64
		invalid []string
	}{
		{`{"vps":{"id":12},"configuration_valid":true,"credit_us":null,"peak_value":null,"peak_is_lower_bound":null}`, nil, []string{}},
		{`{"vps":{"id":12},"configuration_valid":true,"credit_us":0,"measurement":"rtnl_latency","peak_unit":"microseconds","peak_value":0,"peak_is_lower_bound":false}`, new(int64), []string{}},
		{`{"vps":{"id":12},"configuration_valid":true,"credit_us":125,"measurement":"rtnl_latency","peak_unit":"microseconds","peak_value":0,"peak_is_lower_bound":false}`, func() *int64 { v := int64(125); return &v }(), []string{}},
		{`{"vps":{"id":12},"configuration_valid":false,"credit_us":null}`, nil, []string{"rtnl_latency"}},
	}
	for _, tc := range cases {
		api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
			if req.Method != "GET" || req.URL.Path != "/v7.0/vpses/12/evidence" || req.URL.RawQuery != "" || req.Body != nil {
				t.Fatalf("child request: %s %s", req.Method, req.URL)
			}
			if req.Header.Get("X-HaveAPI-Auth-Token") != "fixture-token" {
				t.Fatal("authentication lost")
			}
			invalid, _ := json.Marshal(tc.invalid)
			valid := "true"
			if len(tc.invalid) > 0 {
				valid = "false"
			}
			body := `{"status":true,"response":{"evidence":[` + tc.body + `],"_meta":{"configuration_valid":` + valid + `,"unavailable_signals":` + string(invalid) + `}}}`
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
		})})
		var page struct {
			client.Envelope
			Response struct {
				Evidence []evidenceRow `json:"evidence"`
				Meta     evidenceMeta  `json:"_meta"`
			} `json:"response"`
		}
		err := api.DoQueryStringRequest("/v7.0/vpses/12/evidence", nil, &page)
		if err != nil {
			t.Fatal(err)
		}
		if !page.Status {
			t.Fatalf("API failure: %s %v", page.Message, page.Errors)
		}
		row := page.Response.Evidence[0]
		if row.Vps == nil || row.Vps.Id != 12 {
			t.Fatal("child row identity lost")
		}
		if !reflect.DeepEqual(row.CreditUs, tc.credit) || !reflect.DeepEqual(page.Response.Meta.UnavailableSignals, tc.invalid) {
			t.Fatalf("null/zero/array: %#v %#v", row, page.Response.Meta)
		}
		if tc.credit == nil {
			if row.PeakValue != nil || row.PeakIsLowerBound != nil {
				t.Fatal("unknown peak became zero/false")
			}
		} else if row.PeakValue == nil || *row.PeakValue != 0 || row.PeakIsLowerBound == nil || *row.PeakIsLowerBound {
			t.Fatal("known zero/false peak lost")
		}
		if len(tc.invalid) == 0 && !row.ConfigurationValid {
			t.Fatal("valid-config unavailable credit must remain distinct from zero")
		}
	}
}

func TestEvidenceFleetAndSavedEnvelopeNamespaces(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		if req.Method != "GET" {
			t.Fatal("unexpected method")
		}
		var body string
		switch req.URL.Path {
		case "/v7.0/vps_evidence":
			want := url.Values{"vps_evidence[offset]": {"0"}, "vps_evidence[limit]": {"100"}, "vps_evidence[vps]": {"12"}, "vps_evidence[signal]": {"rtnl_latency"}}
			if !reflect.DeepEqual(req.URL.Query(), want) {
				t.Fatalf("fleet query: %v", req.URL.Query())
			}
			body = `{"status":true,"response":{"vps_evidence":[{"vps":{"id":12},"credit_us":125,"peak_value":null,"peak_is_lower_bound":null}],"_meta":{"configuration_valid":true,"unavailable_signals":[],"total_count":2,"next_offset":1}}}`
		case "/v7.0/incident_reports/12/evidence_intervals":
			if !reflect.DeepEqual(req.URL.Query(), url.Values{"evidence_interval[offset]": {"0"}, "evidence_interval[limit]": {"100"}}) {
				t.Fatalf("saved query: %v", req.URL.Query())
			}
			body = `{"status":true,"response":{"evidence_intervals":[{"credit_us":1530865,"from_at":"2026-11-25T12:00:00.123456Z","to_at":"2026-11-25T12:00:01.654321Z","measurement":null,"peak_unit":null,"peak_value":null,"peak_is_lower_bound":null}],"_meta":{"available":true,"total_count":2,"next_offset":1,"condition_measurement":"rtnl_latency","condition_operator":"gt","condition_value":1000000,"condition_unit":"microseconds","window_complete":false}}}`
		default:
			t.Fatalf("unexpected path %s", req.URL)
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})})
	var fleet struct {
		client.Envelope
		Response struct {
			Rows []evidenceRow `json:"vps_evidence"`
			Meta evidenceMeta  `json:"_meta"`
		} `json:"response"`
	}
	if err := api.DoQueryStringRequest("/v7.0/vps_evidence", map[string]string{"vps_evidence[offset]": "0", "vps_evidence[limit]": "100", "vps_evidence[vps]": "12", "vps_evidence[signal]": "rtnl_latency"}, &fleet); err != nil {
		t.Fatal(err)
	}
	if !fleet.Status || fleet.Response.Meta.TotalCount != 2 || fleet.Response.Meta.NextOffset == nil || *fleet.Response.Meta.NextOffset != 1 || *fleet.Response.Rows[0].CreditUs != 125 || fleet.Response.Rows[0].Vps == nil || fleet.Response.Rows[0].Vps.Id != 12 {
		t.Fatalf("fleet envelope: %#v", fleet)
	}
	var saved struct {
		client.Envelope
		Response struct {
			Rows []evidenceRow `json:"evidence_intervals"`
			Meta evidenceMeta  `json:"_meta"`
		} `json:"response"`
	}
	if err := api.DoQueryStringRequest("/v7.0/incident_reports/12/evidence_intervals", map[string]string{"evidence_interval[offset]": "0", "evidence_interval[limit]": "100"}, &saved); err != nil {
		t.Fatal(err)
	}
	if !saved.Status {
		t.Fatalf("API failure: %s %v", saved.Message, saved.Errors)
	}
	row, meta := saved.Response.Rows[0], saved.Response.Meta
	if *row.FromAt != "2026-11-25T12:00:00.123456Z" || *row.ToAt != "2026-11-25T12:00:01.654321Z" || *row.CreditUs != 1530865 || row.PeakValue != nil || meta.ConditionValue == nil || *meta.ConditionValue != 1000000 || meta.ConditionOperator == nil || *meta.ConditionOperator != "gt" || meta.WindowComplete == nil || *meta.WindowComplete {
		t.Fatalf("saved values: %#v %#v", row, meta)
	}
}

func TestEvidenceLowLevelErrorsAndGeneratedScalarLimit(t *testing.T) {
	var generated client.ActionVpsEvidenceIndexOutput
	if err := json.Unmarshal([]byte(`{"credit_us":null,"peak_value":null,"peak_is_lower_bound":null}`), &generated); err != nil {
		t.Fatal(err)
	}
	if generated.CreditUs != 0 || generated.PeakValue != 0 || generated.PeakIsLowerBound {
		t.Fatal("retained generator scalar behavior changed")
	}
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 403, Body: io.NopCloser(strings.NewReader(`{"status":false,"message":"denied","errors":{"vps":["unavailable"]}}`))}, nil
	})})
	var result struct{ client.Envelope }
	if err := api.DoQueryStringRequest("/v7.0/vpses/12/evidence", nil, &result); err != nil {
		t.Fatal(err)
	}
	if result.Status || result.Message != "denied" || !reflect.DeepEqual(result.Errors["vps"], []string{"unavailable"}) {
		t.Fatalf("API error lost: %#v", result)
	}
	failure := errors.New("transport failure")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) { return nil, failure })})
	if err := api.DoQueryStringRequest("/v7.0/vpses/12/evidence", nil, &result); !errors.Is(err, failure) {
		t.Fatalf("transport error: %v", err)
	}
}

func TestProjectedEvidenceOmitsChildAndKeepsGeneratedFleetRoute(t *testing.T) {
	api := client.New("https://api.example.test")
	if _, exists := reflect.TypeOf(*api.Vps).FieldByName("Evidence"); exists {
		t.Fatal("colliding child accessor must not be generated")
	}
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		if req.Method != "GET" || req.URL.Path != "/v7.0/vps_evidence" {
			t.Fatalf("generated fleet route: %s %s", req.Method, req.URL)
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{"vps_evidence":[{"credit_us":125}]}}`))}, nil
	})})
	page, err := api.VpsEvidence.Index.Prepare().Call()
	if err != nil {
		t.Fatal(err)
	}
	if !page.Status || len(page.Output) != 1 || page.Output[0].CreditUs != 125 {
		t.Fatalf("generated fleet output namespace: %#v", page)
	}
}
