// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestConnectionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "list",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestConnectionsAttach(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "attach",
			"--agent-id", "agent_abc123",
			"--connection-id", "connectionId",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connectionId: connectionId\n" +
			"metadata:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connections", "attach",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestConnectionsDetach(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "detach",
			"--agent-id", "agent_abc123",
			"--connection-id", "env:org_abc123:fleet_abc123::conn_abc123",
		)
	})
}
