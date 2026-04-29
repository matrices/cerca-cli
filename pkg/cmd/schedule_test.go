// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/matrices/cerca-cli/internal/mocktest"
)

func TestSchedulesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"schedules", "create",
			"--agent-id", "agent_abc123",
			"--name", "name",
			"--prompt", "prompt",
			"--cron", "cron",
			"--instructions", "instructions",
			"--model", "model",
			"--run-at", "'2019-12-27T18:11:19.117Z'",
			"--timezone", "timezone",
			"--tool", "sandbox.*",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"prompt: prompt\n" +
			"cron: cron\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"runAt: '2019-12-27T18:11:19.117Z'\n" +
			"timezone: timezone\n" +
			"tools:\n" +
			"  - sandbox.*\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"schedules", "create",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestSchedulesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"schedules", "update",
			"--agent-id", "agent_abc123",
			"--schedule-id", "schedule_abc123",
			"--cron", "cron",
			"--enabled=true",
			"--instructions", "instructions",
			"--model", "model",
			"--name", "name",
			"--prompt", "prompt",
			"--run-at", "'2019-12-27T18:11:19.117Z'",
			"--timezone", "timezone",
			"--tool", "sandbox.*",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cron: cron\n" +
			"enabled: true\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"name: name\n" +
			"prompt: prompt\n" +
			"runAt: '2019-12-27T18:11:19.117Z'\n" +
			"timezone: timezone\n" +
			"tools:\n" +
			"  - sandbox.*\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"schedules", "update",
			"--agent-id", "agent_abc123",
			"--schedule-id", "schedule_abc123",
		)
	})
}

func TestSchedulesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"schedules", "list",
			"--agent-id", "agent_abc123",
		)
	})
}

func TestSchedulesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"schedules", "delete",
			"--agent-id", "agent_abc123",
			"--schedule-id", "schedule_abc123",
		)
	})
}

func TestSchedulesTrigger(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"schedules", "trigger",
			"--agent-id", "agent_abc123",
			"--schedule-id", "schedule_abc123",
		)
	})
}
