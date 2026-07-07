package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		authValue string
		wantKey   string
		wantErr   bool
	}{
		{
			name:      "valid api key",
			authValue: "ApiKey abc123",
			wantKey:   "abc123",
			wantErr:   false,
		},
		{
			name:      "invalid api key",
			authValue: "Api abc123",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.authValue != "" {
				headers.Set("Authorization", tt.authValue)
			}

			got, err := GetAPIKey(headers)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.wantKey {
				t.Fatalf("got %q, want %q", got, tt.wantKey)
			}
		})
	}
}
