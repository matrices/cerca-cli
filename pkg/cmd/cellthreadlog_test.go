// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsThreadsLogsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads:logs", "list",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}
