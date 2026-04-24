// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestToolSourcesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sources", "create",
			"--fleet-id", "fleet_abc123",
			"--source", "{auth: {kind: bar}, namespace: docs, tools: [{name: bar, description: bar, inputSchema: bar, endpoint: bar}], type: http, approval: always, enabled: true, execution: {timeoutMs: bar, maxAttempts: bar, retryMode: bar}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  kind: bar\n" +
			"namespace: docs\n" +
			"tools:\n" +
			"  - name: bar\n" +
			"    description: bar\n" +
			"    inputSchema: bar\n" +
			"    endpoint: bar\n" +
			"type: http\n" +
			"approval: always\n" +
			"enabled: true\n" +
			"execution:\n" +
			"  timeoutMs: bar\n" +
			"  maxAttempts: bar\n" +
			"  retryMode: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sources", "create",
			"--fleet-id", "fleet_abc123",
		)
	})
}

func TestToolSourcesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sources", "retrieve",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestToolSourcesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sources", "update",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
			"--source", "{auth: {kind: bar}, namespace: docs, tools: [{name: bar, description: bar, inputSchema: bar, endpoint: bar}], type: http, approval: always, enabled: true, execution: {timeoutMs: bar, maxAttempts: bar, retryMode: bar}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  kind: bar\n" +
			"namespace: docs\n" +
			"tools:\n" +
			"  - name: bar\n" +
			"    description: bar\n" +
			"    inputSchema: bar\n" +
			"    endpoint: bar\n" +
			"type: http\n" +
			"approval: always\n" +
			"enabled: true\n" +
			"execution:\n" +
			"  timeoutMs: bar\n" +
			"  maxAttempts: bar\n" +
			"  retryMode: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sources", "update",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestToolSourcesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sources", "list",
			"--max-items", "10",
			"--fleet-id", "fleet_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestToolSourcesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sources", "delete",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestToolSourcesListForAgent(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sources", "list-for-agent",
			"--agent-id", "agent_abc123",
		)
	})
}
