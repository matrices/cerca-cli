// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsEventsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:events", "list",
			"--cell-id", "cell_abc123",
			"--cursor", "cursor_abc123",
			"--events", "thread.created,thread.completed",
			"--history", "true",
			"--limit", "20",
		)
	})
}

func TestCellsEventsSubscribe(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:events", "subscribe",
			"--cell-id", "cell_abc123",
		)
	})
}
