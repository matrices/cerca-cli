// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestToolsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tools", "create",
			"--fleet-id", "fleet_abc123",
			"--tool", "{auth: {kind: none}, namespace: docs, tools: [{description: Search documents, endpoint: {method: GET, url: https://docs.example.com/search, body: json_params, headers: {foo: string}, path: {foo: params.query}, query: {foo: params.query}}, inputSchema: {type: bar}, name: search, approval: always, execution: {idempotencyKeyHeader: idempotencyKeyHeader, maxAttempts: 3, retryMode: safe_only, timeoutMs: 10000}, response: {mode: auto}}], type: http, approval: always, enabled: true, execution: {idempotencyKeyHeader: idempotencyKeyHeader, maxAttempts: 3, retryMode: safe_only, timeoutMs: 10000}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  kind: none\n" +
			"namespace: docs\n" +
			"tools:\n" +
			"  - description: Search documents\n" +
			"    endpoint:\n" +
			"      method: GET\n" +
			"      url: https://docs.example.com/search\n" +
			"      body: json_params\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      path:\n" +
			"        foo: params.query\n" +
			"      query:\n" +
			"        foo: params.query\n" +
			"    inputSchema:\n" +
			"      type: bar\n" +
			"    name: search\n" +
			"    approval: always\n" +
			"    execution:\n" +
			"      idempotencyKeyHeader: idempotencyKeyHeader\n" +
			"      maxAttempts: 3\n" +
			"      retryMode: safe_only\n" +
			"      timeoutMs: 10000\n" +
			"    response:\n" +
			"      mode: auto\n" +
			"type: http\n" +
			"approval: always\n" +
			"enabled: true\n" +
			"execution:\n" +
			"  idempotencyKeyHeader: idempotencyKeyHeader\n" +
			"  maxAttempts: 3\n" +
			"  retryMode: safe_only\n" +
			"  timeoutMs: 10000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tools", "create",
			"--fleet-id", "fleet_abc123",
		)
	})
}

func TestToolsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tools", "retrieve",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestToolsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tools", "update",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
			"--tool", "{auth: {kind: none}, namespace: docs, tools: [{description: Search documents, endpoint: {method: GET, url: https://docs.example.com/search, body: json_params, headers: {foo: string}, path: {foo: params.query}, query: {foo: params.query}}, inputSchema: {type: bar}, name: search, approval: always, execution: {idempotencyKeyHeader: idempotencyKeyHeader, maxAttempts: 3, retryMode: safe_only, timeoutMs: 10000}, response: {mode: auto}}], type: http, approval: always, enabled: true, execution: {idempotencyKeyHeader: idempotencyKeyHeader, maxAttempts: 3, retryMode: safe_only, timeoutMs: 10000}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"auth:\n" +
			"  kind: none\n" +
			"namespace: docs\n" +
			"tools:\n" +
			"  - description: Search documents\n" +
			"    endpoint:\n" +
			"      method: GET\n" +
			"      url: https://docs.example.com/search\n" +
			"      body: json_params\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      path:\n" +
			"        foo: params.query\n" +
			"      query:\n" +
			"        foo: params.query\n" +
			"    inputSchema:\n" +
			"      type: bar\n" +
			"    name: search\n" +
			"    approval: always\n" +
			"    execution:\n" +
			"      idempotencyKeyHeader: idempotencyKeyHeader\n" +
			"      maxAttempts: 3\n" +
			"      retryMode: safe_only\n" +
			"      timeoutMs: 10000\n" +
			"    response:\n" +
			"      mode: auto\n" +
			"type: http\n" +
			"approval: always\n" +
			"enabled: true\n" +
			"execution:\n" +
			"  idempotencyKeyHeader: idempotencyKeyHeader\n" +
			"  maxAttempts: 3\n" +
			"  retryMode: safe_only\n" +
			"  timeoutMs: 10000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tools", "update",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}

func TestToolsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tools", "list",
			"--max-items", "10",
			"--fleet-id", "fleet_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestToolsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tools", "delete",
			"--fleet-id", "fleet_abc123",
			"--source-id", "toolsrc_abc123",
		)
	})
}
