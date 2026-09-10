package reducto

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeUnmarshal(t *testing.T) {
	cases := map[string]time.Time{
		`"2026-09-06T05:17:35.492581"`:  time.Date(2026, 9, 6, 5, 17, 35, 492581000, time.UTC),
		`"2026-09-06T05:17:35"`:         time.Date(2026, 9, 6, 5, 17, 35, 0, time.UTC),
		`"2026-09-06T05:17:35Z"`:        time.Date(2026, 9, 6, 5, 17, 35, 0, time.UTC),
		`"2026-09-06T05:17:35.5+02:00"`: time.Date(2026, 9, 6, 5, 17, 35, 500000000, time.FixedZone("", 2*3600)),
	}
	for raw, want := range cases {
		var got Time
		if err := json.Unmarshal([]byte(raw), &got); err != nil {
			t.Fatalf("%s: %v", raw, err)
		}
		if !got.Equal(want) {
			t.Errorf("%s: got %v, want %v", raw, got.Time, want)
		}
	}
}

func TestTimeUnmarshalNullAndInvalid(t *testing.T) {
	var got Time
	if err := json.Unmarshal([]byte("null"), &got); err != nil || !got.IsZero() {
		t.Fatalf("null: err=%v zero=%v", err, got.IsZero())
	}
	if err := json.Unmarshal([]byte(`"yesterday"`), &got); err == nil {
		t.Fatal("expected an error for an unparseable timestamp")
	}
}

func TestTimeRoundTrip(t *testing.T) {
	var job SingleJob
	if err := json.Unmarshal([]byte(`{"job_id":"a","status":"Completed","type":"Parse","raw_config":"{}","created_at":"2024-01-01T00:00:00"}`), &job); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(job.CreatedAt)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `"2024-01-01T00:00:00Z"` {
		t.Errorf("marshal: got %s", out)
	}
}
