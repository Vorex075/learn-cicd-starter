package auth_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		apiKey        string
		header        string
		keyPreamble   string
		errorExpected bool
	}{
		"correct api key in header": {
			apiKey:        "hello123",
			keyPreamble:   "ApiKey",
			header:        "Authorization",
			errorExpected: false,
		},
		"api key set in the inccorrect header": {
			apiKey:        "hello123",
			keyPreamble:   "ApiKey",
			header:        "Not-Auth",
			errorExpected: true,
		},
		"no api key set": {
			apiKey:        "",
			keyPreamble:   "",
			header:        "Authorization",
			errorExpected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			realHeader := http.Header{tc.header: {fmt.Sprintf("%s %s", tc.keyPreamble, tc.apiKey)}}
			t.Logf("Header: %v", realHeader)
			apiKey, err := auth.GetAPIKey(realHeader)

			t.Logf("Error returned?: %v", err != nil)

			if !tc.errorExpected && err == nil && apiKey != tc.apiKey {
				t.Fatalf("In test '%s': result api key differs from original: '%s' - '%s'",
					name, tc.apiKey, apiKey)
			} else if !tc.errorExpected && err != nil {
				t.Fatalf("In test '%s': %v", name, err)
			} else if tc.errorExpected && err == nil {
				t.Fatalf("In test '%s': error expected. None found", name)
			}
		})
	}
}
