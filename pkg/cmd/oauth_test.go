// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestOAuthConnect(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth", "connect",
			"--provider", "google",
			"--return-origin", "https://app.example.com",
			"--scope", "env:org_abc123:fleet_abc123",
			"--scope", "email",
			"--scope", "profile",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"returnOrigin: https://app.example.com\n" +
			"scope: env:org_abc123:fleet_abc123\n" +
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
