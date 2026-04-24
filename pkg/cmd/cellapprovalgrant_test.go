// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsApprovalGrantsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:approval-grants", "list",
			"--cell-id", "cell_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestCellsApprovalGrantsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:approval-grants", "delete",
			"--cell-id", "cell_abc123",
			"--grant-id", "grant_abc123",
		)
	})
}
