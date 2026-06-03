package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{
			input: http.Header{"Authorization": []string{"ApiKey 12345"}},
			want:  "12345",
		},
		{
			input: http.Header{"Authorization": []string{"ApiKey abcde"}},
			want:  "abcde",
		},
		{
			input: http.Header{"Authorization": []string{"Bearer 12345"}},
			want:  "",
		},
		{
			input: http.Header{"Authorization": []string{""}},
			want:  "",
		},
		{
			input: http.Header{},
			want:  "",
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if got != tc.want {
			t.Errorf("GetAPIKey(%v) = %v; want %v", tc.input, got, tc.want)
		}
		if tc.want == "" && err == nil {
			t.Errorf("GetAPIKey(%v) expected an error but got none", tc.input)
		}
	}
}
