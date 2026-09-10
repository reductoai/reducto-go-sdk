package reducto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// WebhookTolerance is how far a webhook timestamp may drift from the local clock.
const WebhookTolerance = 5 * time.Minute

// WebhookVerificationError is returned by VerifyWebhook when the delivery cannot be trusted.
type WebhookVerificationError struct {
	Reason string
}

func (e *WebhookVerificationError) Error() string {
	return "reducto: webhook verification failed: " + e.Reason
}

// VerifyWebhook checks the signature of a webhook delivered in Svix mode. secret is the endpoint
// secret from the webhook portal (with or without the "whsec_" prefix). payload is the raw
// request body; header is the incoming request's headers. Both the svix-* and webhook-* header
// names are accepted. Deliveries in "direct" mode are unsigned and cannot be verified.
func VerifyWebhook(secret string, payload []byte, header http.Header) error {
	return verifyWebhook(secret, payload, header, time.Now())
}

func verifyWebhook(secret string, payload []byte, header http.Header, now time.Time) error {
	id := firstHeader(header, "svix-id", "webhook-id")
	ts := firstHeader(header, "svix-timestamp", "webhook-timestamp")
	sig := firstHeader(header, "svix-signature", "webhook-signature")
	if id == "" || ts == "" || sig == "" {
		return &WebhookVerificationError{Reason: "missing svix-id, svix-timestamp or svix-signature header"}
	}
	unix, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return &WebhookVerificationError{Reason: "malformed timestamp"}
	}
	if drift := now.Sub(time.Unix(unix, 0)); drift > WebhookTolerance || drift < -WebhookTolerance {
		return &WebhookVerificationError{Reason: "timestamp outside tolerance"}
	}
	key, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(secret, "whsec_"))
	if err != nil {
		return &WebhookVerificationError{Reason: "secret is not base64"}
	}
	mac := hmac.New(sha256.New, key)
	fmt.Fprintf(mac, "%s.%s.", id, ts)
	mac.Write(payload)
	want := mac.Sum(nil)
	for _, entry := range strings.Fields(sig) {
		version, value, ok := strings.Cut(entry, ",")
		if !ok || version != "v1" {
			continue
		}
		got, err := base64.StdEncoding.DecodeString(value)
		if err == nil && hmac.Equal(got, want) {
			return nil
		}
	}
	return &WebhookVerificationError{Reason: "no matching signature"}
}

func firstHeader(h http.Header, names ...string) string {
	for _, n := range names {
		if v := h.Get(n); v != "" {
			return v
		}
	}
	return ""
}
