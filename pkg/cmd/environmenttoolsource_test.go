// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestEnvironmentsToolSourcesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:tool-sources", "create",
			"--environment-id", "env_abc123",
			"--auth", "{kind: none}",
			"--namespace", "docs",
			"--tool", "{name: search, description: Search documents, inputSchema: {type: object}, endpoint: {method: GET, url: https://docs.example.com/search}}",
			"--type", "http",
			"--approval", "always",
			"--enabled=true",
			"--execution", "{timeoutMs: 10000, maxAttempts: 3, retryMode: safe_only}",
			"--url", "https://mcp.example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  kind: none\n" +
			"namespace: docs\n" +
			"tools:\n" +
			"  - name: search\n" +
			"    description: Search documents\n" +
			"    inputSchema:\n" +
			"      type: object\n" +
			"    endpoint:\n" +
			"      method: GET\n" +
			"      url: https://docs.example.com/search\n" +
			"type: http\n" +
			"approval: always\n" +
			"enabled: true\n" +
			"execution:\n" +
			"  timeoutMs: 10000\n" +
			"  maxAttempts: 3\n" +
			"  retryMode: safe_only\n" +
			"url: https://mcp.example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"environments:tool-sources", "create",
			"--environment-id", "env_abc123",
		)
	})
}

func TestEnvironmentsToolSourcesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:tool-sources", "retrieve",
			"--environment-id", "env_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestEnvironmentsToolSourcesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:tool-sources", "update",
			"--environment-id", "env_abc123",
			"--source-id", "toolsrc_abc123",
			"--auth", "{kind: none}",
			"--namespace", "docs",
			"--tool", "{name: search, description: Search documents, inputSchema: {type: object}, endpoint: {method: GET, url: https://docs.example.com/search}}",
			"--type", "http",
			"--approval", "always",
			"--enabled=true",
			"--execution", "{timeoutMs: 10000, maxAttempts: 3, retryMode: safe_only}",
			"--url", "https://mcp.example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  kind: none\n" +
			"namespace: docs\n" +
			"tools:\n" +
			"  - name: search\n" +
			"    description: Search documents\n" +
			"    inputSchema:\n" +
			"      type: object\n" +
			"    endpoint:\n" +
			"      method: GET\n" +
			"      url: https://docs.example.com/search\n" +
			"type: http\n" +
			"approval: always\n" +
			"enabled: true\n" +
			"execution:\n" +
			"  timeoutMs: 10000\n" +
			"  maxAttempts: 3\n" +
			"  retryMode: safe_only\n" +
			"url: https://mcp.example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"environments:tool-sources", "update",
			"--environment-id", "env_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestEnvironmentsToolSourcesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:tool-sources", "list",
			"--environment-id", "env_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestEnvironmentsToolSourcesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:tool-sources", "delete",
			"--environment-id", "env_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}
