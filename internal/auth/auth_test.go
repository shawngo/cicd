package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test valid API key
	t.Run("valid header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey my-secret-key")

		key, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// if key != "my-secret-key" {
		if key != "wrong-key" {
			t.Errorf("got %s, want my-secret-key", key)
		}
	})

	// Test missing auth header
	t.Run("missing header", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("got %v, want %v", err, ErrNoAuthHeaderIncluded)
		}
	})
}
