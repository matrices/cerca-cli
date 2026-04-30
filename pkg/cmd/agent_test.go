// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
	"github.com/matrices/cerca-cli/internal/requestflag"
)

func TestAgentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--configuration", "{approvals: {timeoutMs: 0, tools: {foo: always}}, defaultModel: defaultModel, instructions: instructions, tools: [sandbox.*]}",
			"--fleet-id", "fleetId",
			"--metadata", "{project: alpha}",
			"--user-id", "userId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--configuration.approvals", "{timeoutMs: 0, tools: {foo: always}}",
			"--configuration.default-model", "defaultModel",
			"--configuration.instructions", "instructions",
			"--configuration.tools", "[sandbox.*]",
			"--fleet-id", "fleetId",
			"--metadata", "{project: alpha}",
			"--user-id", "userId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"configuration:\n" +
			"  approvals:\n" +
			"    timeoutMs: 0\n" +
			"    tools:\n" +
			"      foo: always\n" +
			"  defaultModel: defaultModel\n" +
			"  instructions: instructions\n" +
			"  tools:\n" +
			"    - sandbox.*\n" +
			"fleetId: fleetId\n" +
			"metadata:\n" +
			"  project: alpha\n" +
			"userId: userId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "create",
		)
	})
}

func TestAgentsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "retrieve",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestAgentsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "agent_abc123",
			"--configuration", "{approvals: {timeoutMs: 0, tools: {foo: always}}, defaultModel: defaultModel, instructions: instructions, tools: [sandbox.*]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "agent_abc123",
			"--configuration.approvals", "{timeoutMs: 0, tools: {foo: always}}",
			"--configuration.default-model", "defaultModel",
			"--configuration.instructions", "instructions",
			"--configuration.tools", "[sandbox.*]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"configuration:\n" +
			"  approvals:\n" +
			"    timeoutMs: 0\n" +
			"    tools:\n" +
			"      foo: always\n" +
			"  defaultModel: defaultModel\n" +
			"  instructions: instructions\n" +
			"  tools:\n" +
			"    - sandbox.*\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestAgentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list",
			"--max-items", "10",
			"--active", "true",
			"--cursor", "cursor_abc123",
			"--fleet-id", "fleet_abc123",
			"--limit", "20",
		)
	})
}

func TestAgentsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "delete",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestAgentsListTools(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list-tools",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestAgentsRetrieveConfig(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "retrieve-config",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestAgentsUpdateMetadata(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update-metadata",
			"--agent-id", "agent_abc123",
			"--metadata", "{project: alpha}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  project: alpha\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "update-metadata",
			"--agent-id", "agent_abc123",
		)
	})
}
