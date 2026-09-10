package reducto

import (
	"bytes"
	"encoding/json"
)

func jsonKind(b []byte) byte {
	b = bytes.TrimLeft(b, " \t\r\n")
	if len(b) == 0 {
		return 0
	}
	switch b[0] {
	case '"':
		return 's'
	case '{':
		return 'o'
	case '[':
		return 'a'
	case 'n':
		return 'n'
	case 't', 'f':
		return 'b'
	}
	return 'd'
}

func isIntegerLiteral(b []byte) bool {
	return !bytes.ContainsAny(b, ".eE")
}

func jsonDiscriminator(b []byte, prop string) string {
	var m map[string]json.RawMessage
	if json.Unmarshal(b, &m) != nil {
		return ""
	}
	var v string
	_ = json.Unmarshal(m[prop], &v)
	return v
}

func jsonHasKeys(b []byte, keys ...string) bool {
	var m map[string]json.RawMessage
	if json.Unmarshal(b, &m) != nil {
		return false
	}
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

func strictUnmarshal[T any](b []byte) (*T, bool) {
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	var v T
	if err := dec.Decode(&v); err != nil {
		return nil, false
	}
	return &v, true
}

func jsonHead(b []byte) string {
	b = bytes.TrimSpace(b)
	if len(b) > 120 {
		return string(b[:120]) + "…"
	}
	return string(b)
}
