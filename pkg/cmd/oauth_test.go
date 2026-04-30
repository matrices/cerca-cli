// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestOAuthConnect(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth", "connect",
			"--provider", "google",
			"--owner", "{type: organization}",
			"--return-origin", "https://app.example.com",
			"--scope", "email",
			"--scope", "profile",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"owner:\n" +
			"  type: organization\n" +
			"returnOrigin: https://app.example.com\n" +
			"scopes:\n" +
			"  - email\n" +
			"  - profile\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"oauth", "connect",
			"--provider", "google",
		)
	})
}
