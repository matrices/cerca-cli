// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestApprovalGrantsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"approval-grants", "list",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestApprovalGrantsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"approval-grants", "delete",
			"--agent-id", "agent_abc123",
			"--grant-id", "grant_abc123",
		)
	})
}

func TestApprovalGrantsDeleteForThread(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"approval-grants", "delete-for-thread",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--grant-id", "grant_abc123",
		)
	})
}

func TestApprovalGrantsListForThread(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"approval-grants", "list-for-thread",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}
