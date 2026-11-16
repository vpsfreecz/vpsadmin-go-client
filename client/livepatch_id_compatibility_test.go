package client

import (
	"encoding/json"
	"testing"
)

func TestLivepatchIDOutputCompatibility(t *testing.T) {
	decoders := []struct {
		name   string
		decode func([]byte) (string, error)
	}{
		{
			name: "node kernel history index",
			decode: func(data []byte) (string, error) {
				var output ActionNodeKernelHistoryIndexOutput
				err := json.Unmarshal(data, &output)
				return output.LivepatchId, err
			},
		},
		{
			name: "node kernel event index",
			decode: func(data []byte) (string, error) {
				var output ActionNodeKernelEventIndexOutput
				err := json.Unmarshal(data, &output)
				return output.LivepatchId, err
			},
		},
		{
			name: "node kernel event show",
			decode: func(data []byte) (string, error) {
				var output ActionNodeKernelEventShowOutput
				err := json.Unmarshal(data, &output)
				return output.LivepatchId, err
			},
		},
	}

	for _, decoder := range decoders {
		t.Run(decoder.name, func(t *testing.T) {
			for _, data := range []string{`{}`, `{"livepatch_id":null}`} {
				value, err := decoder.decode([]byte(data))
				if err != nil {
					t.Fatal(err)
				}
				if value != "" {
					t.Fatalf("expected empty value for %s, got %q", data, value)
				}
			}

			value, err := decoder.decode([]byte(`{"livepatch_id":"lp_test"}`))
			if err != nil {
				t.Fatal(err)
			}
			if value != "lp_test" {
				t.Fatalf("expected decoded livepatch ID, got %q", value)
			}
		})
	}
}
