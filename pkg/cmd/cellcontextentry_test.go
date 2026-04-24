// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsContextEntriesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:context:entries", "retrieve",
			"--cell-id", "cell_abc123",
			"--key", "notes/project.md",
		)
	})
}

func TestCellsContextEntriesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:context:entries", "delete",
			"--cell-id", "cell_abc123",
			"--key", "notes/project.md",
		)
	})
}

func TestCellsContextEntriesPut(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:context:entries", "put",
			"--cell-id", "cell_abc123",
			"--content", "content",
			"--key", "key",
			"--expected-version", "0",
			"--mime-type", "mimeType",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content: content\n" +
			"key: key\n" +
			"expectedVersion: 0\n" +
			"mimeType: mimeType\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:context:entries", "put",
			"--cell-id", "cell_abc123",
		)
	})
}
