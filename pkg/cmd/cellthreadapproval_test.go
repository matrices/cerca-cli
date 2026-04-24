// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsThreadsApprovalsResolve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads:approvals", "resolve",
			"--cell-id", "cell_abc123",
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
			"cells:threads:approvals", "resolve",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
			"--approval-id", "approval_abc123",
		)
	})
}
