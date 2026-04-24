// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestApprovalsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"approvals", "list",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestApprovalsResolve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"approvals", "resolve",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--approval-id", "approval_abc123",
			"--decision", "approve",
			"--grant", "thread",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"decision: approve\n" +
			"grant: thread\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"approvals", "resolve",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--approval-id", "approval_abc123",
		)
	})
}
