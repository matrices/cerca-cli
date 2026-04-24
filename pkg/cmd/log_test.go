// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestLogsListForAgent(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"logs", "list-for-agent",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestLogsListForThread(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"logs", "list-for-thread",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}
