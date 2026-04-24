// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestEnvironmentsEventsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:events", "list",
			"--env-id", "env_abc123",
			"--cursor", "cursor_abc123",
			"--events", "thread.created,thread.completed",
			"--history", "true",
			"--limit", "20",
		)
	})
}

func TestEnvironmentsEventsSubscribe(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:events", "subscribe",
			"--env-id", "env_abc123",
		)
	})
}
