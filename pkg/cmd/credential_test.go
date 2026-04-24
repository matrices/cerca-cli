// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestCredentialsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"credentials", "list",
			"--max-items", "10",
			"--scope", "env:org_abc123:fleet_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestCredentialsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"credentials", "delete",
			"--scope", "env:org_abc123:fleet_abc123",
			"--connection-id", "env:org_abc123:fleet_abc123::conn_abc123",
		)
	})
}

func TestCredentialsCreateAPIKey(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"credentials", "create-api-key",
			"--scope", "env:org_abc123:fleet_abc123",
			"--api-key", "sk_live_...",
			"--provider", "custom",
			"--account-label", "primary",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"apiKey: sk_live_...\n" +
			"provider: custom\n" +
			"accountLabel: primary\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"credentials", "create-api-key",
			"--scope", "env:org_abc123:fleet_abc123",
		)
	})
}
