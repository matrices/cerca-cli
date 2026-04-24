// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestEventsListForAgent(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "list-for-agent",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--cursor", "cursor_abc123",
			"--events", "thread.created,thread.completed",
			"--history", "true",
			"--limit", "20",
		)
	})
}

func TestEventsListForFleet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "list-for-fleet",
			"--max-items", "10",
			"--fleet-id", "fleet_abc123",
			"--cursor", "cursor_abc123",
			"--events", "thread.created,thread.completed",
			"--history", "true",
			"--limit", "20",
		)
	})
}

func TestEventsListForThread(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "list-for-thread",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--cursor", "cursor_abc123",
			"--events", "thread.created,thread.completed",
			"--history", "true",
			"--limit", "20",
		)
	})
}

func TestEventsSubscribeAgent(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "subscribe-agent",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestEventsSubscribeFleet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "subscribe-fleet",
			"--fleet-id", "fleet_abc123",
		)
	})
}

func TestEventsSubscribeThread(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"events", "subscribe-thread",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}
