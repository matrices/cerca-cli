// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestContextRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"context", "retrieve",
			"--agent-id", "agent_abc123",
			"--key", "notes/project.md",
		)
	})
}

func TestContextList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"context", "list",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
			"--prefix", "notes/",
		)
	})
}

func TestContextDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"context", "delete",
			"--agent-id", "agent_abc123",
			"--key", "notes/project.md",
		)
	})
}

func TestContextSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"context", "search",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--q", "project notes",
			"--cursor", "cursor_abc123",
			"--limit", "20",
			"--prefix", "notes/",
		)
	})
}

func TestContextWrite(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"context", "write",
			"--agent-id", "agent_abc123",
			"--content", "content",
			"--key", "key",
			"--expected-version", "0",
			"--mime-type", "mimeType",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"key: key\n" +
			"expectedVersion: 0\n" +
			"mimeType: mimeType\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"context", "write",
			"--agent-id", "agent_abc123",
		)
	})
}
