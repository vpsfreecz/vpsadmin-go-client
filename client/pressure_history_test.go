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

func TestPressureHistoryAndFleetFilterQueries(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
		if req.Method != "GET" || req.Body != nil {
			t.Fatalf("read method/body: %s", req.Method)
		}
		var want url.Values
		namespace := "vps_pressure_history"
		switch req.URL.Path {
		case "/v7.0/vps_pressure_history":
			want = url.Values{"vps_pressure_history[vps]": {"12"}}
			if _, present := req.URL.Query()["vps_pressure_history[days]"]; present {
				want.Set("vps_pressure_history[days]", "30")
			}
		case "/v7.0/vps_evidence":
			namespace = "vps_evidence"
			want = url.Values{"vps_evidence[user]": {"1"}, "vps_evidence[node]": {"2"}, "vps_evidence[location]": {"3"}, "vps_evidence[environment]": {"4"}, "vps_evidence[offset]": {"0"}, "vps_evidence[limit]": {"25"}}
		default:
			t.Fatalf("unexpected route: %s", req.URL)
		}
		if !reflect.DeepEqual(req.URL.Query(), want) {
			t.Fatalf("namespaced inputs: %v; want %v", req.URL.Query(), want)
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"status":true,"response":{"` + namespace + `":[]}}`))}, nil
	})})
	for _, days := range []bool{false, true} {
		action := api.VpsPressureHistory.Index.Prepare()
		action.NewInput().SetVps(12)
		if days {
			action.Input.SetDays(30)
		}
		page, err := action.Call()
		if err != nil || !page.Status || len(page.Output) != 0 {
			t.Fatalf("history read: %#v %v", page, err)
		}
	}
	action := api.VpsEvidence.Index.Prepare()
	action.NewInput().SetUser(1).SetNode(2).SetLocation(3).SetEnvironment(4).SetOffset(0).SetLimit(25)
	if _, err := action.Call(); err != nil {
		t.Fatal(err)
	}
}

func TestPressureHistoryCustomEnvelopePreservesNullMicrosecondsAndEpoch(t *testing.T) {
	api := client.New("https://api.example.test")
	stamp := "2026-11-28T12:00:00.123456Z"
	for _, known := range []bool{false, true} {
		var value interface{}
		if known {
			value = int64(125)
		}
		body, _ := json.Marshal(map[string]interface{}{"status": true, "response": map[string]interface{}{
			"vps_pressure_history": []interface{}{
				map[string]interface{}{"id": "2026-11-27", "from_at": "2026-11-27T00:00:00.000000Z", "to_at": "2026-11-28T00:00:00.000000Z", "recorded_us": nil, "partial_day": false},
				map[string]interface{}{"id": "2026-11-28", "from_at": "2026-11-28T00:00:00.000000Z", "to_at": "2026-11-29T00:00:00.000000Z", "recorded_us": value, "partial_day": true}},
			"_meta": map[string]interface{}{"available": true, "accounting_since": nil, "ownership_limited": false, "read_at": stamp, "timezone": "UTC", "total_recorded_us": value, "first_recorded_at": nil}}})
		api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(req *http.Request) (*http.Response, error) {
			if req.URL.Path != "/v7.0/vps_pressure_history" || req.URL.Query().Get("vps_pressure_history[vps]") != "12" {
				t.Fatalf("custom route: %s", req.URL)
			}
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(string(body)))}, nil
		})})
		var page struct {
			client.Envelope
			Response struct {
				Rows []struct {
					RecordedUs *int64 `json:"recorded_us"`
					PartialDay *bool  `json:"partial_day"`
				} `json:"vps_pressure_history"`
				Meta struct {
					Available       bool    `json:"available"`
					TotalRecordedUs *int64  `json:"total_recorded_us"`
					AccountingSince *string `json:"accounting_since"`
					FirstRecordedAt *string `json:"first_recorded_at"`
					ReadAt          string  `json:"read_at"`
				} `json:"_meta"`
			} `json:"response"`
		}
		if err := api.DoQueryStringRequest("/v7.0/vps_pressure_history", map[string]string{"vps_pressure_history[vps]": "12"}, &page); err != nil {
			t.Fatal(err)
		}
		rows, meta := page.Response.Rows, page.Response.Meta
		if !page.Status || !meta.Available || len(rows) != 2 || rows[0].RecordedUs != nil || *rows[0].PartialDay || !*rows[1].PartialDay || meta.AccountingSince != nil || meta.FirstRecordedAt != nil || meta.ReadAt != stamp {
			t.Fatalf("history facts: %#v", page)
		}
		if known {
			if rows[1].RecordedUs == nil || *rows[1].RecordedUs != 125 || meta.TotalRecordedUs == nil || *meta.TotalRecordedUs != 125 {
				t.Fatal("subsecond duration lost")
			}
		} else if rows[1].RecordedUs != nil || meta.TotalRecordedUs != nil {
			t.Fatal("unknown history became measured zero")
		}
	}
}
