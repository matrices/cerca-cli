// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestConnectionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "create",
			"--api-key", "sk_live_...",
			"--owner", "{type: organization}",
			"--provider", "custom",
			"--account-label", "primary",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"apiKey: sk_live_...\n" +
			"owner:\n" +
			"  type: organization\n" +
			"provider: custom\n" +
			"accountLabel: primary\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connections", "create",
		)
	})
}

func TestConnectionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "list",
			"--max-items", "10",
			"--owner-type", "fleet",
			"--cursor", "cursor_abc123",
			"--fleet-id", "fleet_abc123",
			"--limit", "20",
		)
	})
}

func TestConnectionsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "delete",
			"--connection-id", "9f063c59-2775-4614-8a68-22e5f90f92f3",
			"--owner-type", "fleet",
			"--fleet-id", "fleet_abc123",
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
			"--connection-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connectionId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
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
			"--connection-id", "9f063c59-2775-4614-8a68-22e5f90f92f3",
		)
	})
}

func TestConnectionsListForAgent(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connections", "list-for-agent",
			"--agent-id", "agent_abc123",
		)
	})
}
