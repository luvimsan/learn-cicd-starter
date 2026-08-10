package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKey -
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "No Auth header",
			headers: http.Header{},
			want: "",
			wantErr: true,
		},
		{
			name: "Valid auth header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			want: "my-secret-key",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		got, err := GetAPIKey(tt.headers)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetApiKey() error = %v, wantErr = %v\n", err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
		}
	}

}
