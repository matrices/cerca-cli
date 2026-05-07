// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestThreadsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "create",
			"--agent-id", "agent_abc123",
			"--instructions", "instructions",
			"--message", "message",
			"--model", "model",
			"--system-prompt", "systemPrompt",
			"--tool", "sandbox.*",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"message: message\n" +
			"model: model\n" +
			"systemPrompt: systemPrompt\n" +
			"tools:\n" +
			"  - sandbox.*\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"threads", "create",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestThreadsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "retrieve",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--debug", "false",
			"--include-messages", "true",
		)
	})
}

func TestThreadsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "list",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
			"--schedule-id", "schedule_abc123",
			"--status", "idle",
		)
	})
}

func TestThreadsActivity(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "activity",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--fleet-id", "fleetId",
		)
	})
}

func TestThreadsCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "cancel",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestThreadsClose(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "close",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestThreadsCompact(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "compact",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestThreadsListMessages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "list-messages",
			"--max-items", "10",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--cursor", "42",
			"--fleet-id", "fleetId",
			"--limit", "100",
		)
	})
}

func TestThreadsStartTurn(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "start-turn",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--message", "message",
			"--model", "model",
			"--tool", "sandbox.*",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"message: message\n" +
			"model: model\n" +
			"tools:\n" +
			"  - sandbox.*\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"threads", "start-turn",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestThreadsSteer(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"threads", "steer",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
			"--message", "message",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("message: message")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"threads", "steer",
			"--agent-id", "agent_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}
