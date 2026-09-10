package reducto

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Time is a time.Time that also decodes the zone-less timestamps the API sends, such as
// "2026-09-06T05:17:35.492581". Those are read as UTC. It marshals as RFC 3339.
type Time struct{ time.Time }

var timeLayouts = []string{time.RFC3339Nano, "2006-01-02T15:04:05"}

func (t *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*t = Time{}
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("reducto: Time: %w", err)
	}
	s = strings.TrimSpace(s)
	if s == "" {
		*t = Time{}
		return nil
	}
	for _, layout := range timeLayouts {
		if parsed, err := time.Parse(layout, s); err == nil {
			*t = Time{parsed}
			return nil
		}
	}
	return fmt.Errorf("reducto: Time: cannot parse %q", s)
}
