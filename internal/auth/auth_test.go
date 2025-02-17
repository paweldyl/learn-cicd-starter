package auth

import (
	"net/http"
	"testing"
)

// GetAPIKey -
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want  string
		err   error
	}{
		{input: http.Header{"Authorization": {"ApiKey testkey"}}, want: "testkey", err: nil},
	}
	for _, test := range tests {
		res, err := GetAPIKey(test.input)
		if res != test.want || err != test.err {
			t.Fatalf(res, err)
		}
	}
}
