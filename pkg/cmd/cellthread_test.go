// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsThreadsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "create",
			"--cell-id", "cell_abc123",
			"--feature", "memory",
			"--instructions", "instructions",
			"--model", "model",
			"--system-prompt", "systemPrompt",
			"--user-message", "userMessage",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"features:\n" +
			"  - memory\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"systemPrompt: systemPrompt\n" +
			"userMessage: userMessage\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:threads", "create",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsThreadsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "retrieve",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
			"--debug", "false",
			"--include-messages", "true",
		)
	})
}

func TestCellsThreadsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "list",
			"--cell-id", "cell_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
			"--schedule-id", "schedule_abc123",
			"--status", "idle",
		)
	})
}

func TestCellsThreadsCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "cancel",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestCellsThreadsClose(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "close",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestCellsThreadsCompact(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "compact",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestCellsThreadsSteer(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "steer",
			"--cell-id", "cell_abc123",
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
			"cells:threads", "steer",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}

func TestCellsThreadsTurn(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:threads", "turn",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
			"--user-message", "userMessage",
			"--feature", "memory",
			"--model", "model",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"userMessage: userMessage\n" +
			"features:\n" +
			"  - memory\n" +
			"model: model\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:threads", "turn",
			"--cell-id", "cell_abc123",
			"--thread-id", "thread_abc123",
		)
	})
}
