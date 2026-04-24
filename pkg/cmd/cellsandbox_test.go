// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsSandboxExec(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:sandbox", "exec",
			"--cell-id", "cell_abc123",
			"--command", "command",
			"--max-buffer", "0",
			"--timeout", "30",
			"--workdir", "/home/user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"command: command\n" +
			"maxBuffer: 0\n" +
			"timeout: 30\n" +
			"workdir: /home/user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:sandbox", "exec",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsSandboxRead(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:sandbox", "read",
			"--cell-id", "cell_abc123",
			"--path", "path",
			"--limit", "0",
			"--offset", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"path: path\n" +
			"limit: 0\n" +
			"offset: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:sandbox", "read",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsSandboxWrite(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:sandbox", "write",
			"--cell-id", "cell_abc123",
			"--content", "content",
			"--path", "path",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"path: path\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:sandbox", "write",
			"--cell-id", "cell_abc123",
		)
	})
}
