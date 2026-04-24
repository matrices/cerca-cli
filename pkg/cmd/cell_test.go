// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
	"github.com/stainless-sdks/cerca-cli/internal/requestflag"
)

func TestCellsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "create",
			"--user-id", "userId",
			"--configuration", "{approvalRequired: [web_search], approvalTimeoutMs: 0, defaultModel: defaultModel, excludedToolSourceIds: [string], features: [memory], instructions: instructions, toolApprovalOverrides: {foo: always}, toolSourceMode: inherit, userContext: userContext}",
			"--environment-id", "environmentId",
			"--metadata", "{project: alpha}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cellsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "create",
			"--user-id", "userId",
			"--configuration.approval-required", "[web_search]",
			"--configuration.approval-timeout-ms", "0",
			"--configuration.default-model", "defaultModel",
			"--configuration.excluded-tool-source-ids", "[string]",
			"--configuration.features", "[memory]",
			"--configuration.instructions", "instructions",
			"--configuration.tool-approval-overrides", "{foo: always}",
			"--configuration.tool-source-mode", "inherit",
			"--configuration.user-context", "userContext",
			"--environment-id", "environmentId",
			"--metadata", "{project: alpha}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"userId: userId\n" +
			"configuration:\n" +
			"  approvalRequired:\n" +
			"    - web_search\n" +
			"  approvalTimeoutMs: 0\n" +
			"  defaultModel: defaultModel\n" +
			"  excludedToolSourceIds:\n" +
			"    - string\n" +
			"  features:\n" +
			"    - memory\n" +
			"  instructions: instructions\n" +
			"  toolApprovalOverrides:\n" +
			"    foo: always\n" +
			"  toolSourceMode: inherit\n" +
			"  userContext: userContext\n" +
			"environmentId: environmentId\n" +
			"metadata:\n" +
			"  project: alpha\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells", "create",
		)
	})
}

func TestCellsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "retrieve",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "update",
			"--cell-id", "cell_abc123",
			"--configuration", "{approvalRequired: [web_search], approvalTimeoutMs: 0, defaultModel: defaultModel, excludedToolSourceIds: [string], features: [memory], instructions: instructions, toolApprovalOverrides: {foo: always}, toolSourceMode: inherit, userContext: userContext}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cellsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "update",
			"--cell-id", "cell_abc123",
			"--configuration.approval-required", "[web_search]",
			"--configuration.approval-timeout-ms", "0",
			"--configuration.default-model", "defaultModel",
			"--configuration.excluded-tool-source-ids", "[string]",
			"--configuration.features", "[memory]",
			"--configuration.instructions", "instructions",
			"--configuration.tool-approval-overrides", "{foo: always}",
			"--configuration.tool-source-mode", "inherit",
			"--configuration.user-context", "userContext",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"configuration:\n" +
			"  approvalRequired:\n" +
			"    - web_search\n" +
			"  approvalTimeoutMs: 0\n" +
			"  defaultModel: defaultModel\n" +
			"  excludedToolSourceIds:\n" +
			"    - string\n" +
			"  features:\n" +
			"    - memory\n" +
			"  instructions: instructions\n" +
			"  toolApprovalOverrides:\n" +
			"    foo: always\n" +
			"  toolSourceMode: inherit\n" +
			"  userContext: userContext\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells", "update",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "list",
			"--active", "true",
			"--cursor", "cursor_abc123",
			"--environment-id", "env_abc123",
			"--limit", "20",
		)
	})
}

func TestCellsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells", "delete",
			"--cell-id", "cell_abc123",
		)
	})
}
