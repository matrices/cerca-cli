// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsConnectionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:connections", "list",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsConnectionsAttach(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:connections", "attach",
			"--cell-id", "cell_abc123",
			"--connection-id", "connectionId",
			"--metadata", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connectionId: connectionId\n" +
			"metadata:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:connections", "attach",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsConnectionsDetach(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:connections", "detach",
			"--cell-id", "cell_abc123",
			"--connection-id", "env:org_abc123:env_abc123::conn_abc123",
		)
	})
}
