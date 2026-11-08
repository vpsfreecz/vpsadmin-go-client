package client_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func TestDnsPrimaryReadinessWireResponse(t *testing.T) {
	api := client.New("https://api.example.test")
	api.SetHTTPClient(&http.Client{Transport: oauthTransport(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodGet || request.URL.Path != "/v7.0/dns_server_zone_primary_transfer_states" {
			t.Fatalf("unexpected path-state request: %s %s", request.Method, request.URL)
		}
		if query := request.URL.Query(); query.Get("dns_server_zone_primary_transfer_state[dns_zone_transfer]") != "20" || query.Get("dns_server_zone_primary_transfer_state[limit]") != "2" {
			t.Fatalf("unexpected path-state filters: %v", query)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{
			"status": true,
			"response": {"dns_server_zone_primary_transfer_states": [
				{"id": 1, "status": "success", "evidence_current": false, "last_attempt_kind": "ixfr_probe", "primary_serial": 100, "secondary_serial": 100},
				{"id": 2, "status": "failed", "evidence_current": true, "failure_class": "network", "reason_code": "timeout"}
			]}
		}`))}, nil
	})})
	action := api.DnsServerZonePrimaryTransferState.Index.Prepare()
	action.NewInput().SetDnsZoneTransfer(20).SetLimit(2)
	response, err := action.Call()
	if err != nil || !response.Status || len(response.Output) != 2 {
		t.Fatalf("unexpected path-state response: %v, %v", response, err)
	}
	if first := response.Output[0]; first.EvidenceCurrent || first.Status != "success" || first.LastAttemptKind != "ixfr_probe" {
		t.Fatalf("stale positive evidence lost its independent currentness: %+v", first)
	}
	if second := response.Output[1]; !second.EvidenceCurrent || second.FailureClass != "network" || second.ReasonCode != "timeout" {
		t.Fatalf("current network failure was not decoded: %+v", second)
	}
}

func TestDnsLegacyLogHasNoInferredPrimary(t *testing.T) {
	var log client.ActionDnsServerZoneTransferLogShowOutput
	if err := json.Unmarshal([]byte(`{"attempt_kind":"legacy","status":"failed","dns_zone_transfer":null,"primary_serial":null}`), &log); err != nil {
		t.Fatal(err)
	}
	if log.AttemptKind != "legacy" || log.DnsZoneTransfer != nil {
		t.Fatalf("historical log gained a current primary: %+v", log)
	}
	var old client.ActionDnsZoneTransferShowOutput
	if err := json.Unmarshal([]byte(`{"id":20,"peer_type":"primary_type"}`), &old); err != nil {
		t.Fatal(err)
	}
	if old.Id != 20 || old.TransferCheckStatus != "" {
		t.Fatalf("old API response invented readiness: %+v", old)
	}
}
