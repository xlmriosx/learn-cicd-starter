package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my_secret_key"},
			},
			expectedKey:   "my_secret_key",
			expectedError: nil,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Header - Wrong Scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer some_token"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Header - Missing Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if key != tc.expectedKey {
				t.Errorf("Expected key: %v, got: %v", tc.expectedKey, key)
			}

			if tc.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error: %v, got nil", tc.expectedError)
				} else if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			}
		})
	}
}
