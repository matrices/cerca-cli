// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestCellsSchedulesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:schedules", "create",
			"--cell-id", "cell_abc123",
			"--cron", "cron",
			"--name", "name",
			"--prompt", "prompt",
			"--feature", "memory",
			"--instructions", "instructions",
			"--model", "model",
			"--timezone", "timezone",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cron: cron\n" +
			"name: name\n" +
			"prompt: prompt\n" +
			"features:\n" +
			"  - memory\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"timezone: timezone\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:schedules", "create",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsSchedulesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:schedules", "update",
			"--cell-id", "cell_abc123",
			"--schedule-id", "schedule_abc123",
			"--cron", "cron",
			"--enabled=true",
			"--feature", "memory",
			"--instructions", "instructions",
			"--model", "model",
			"--name", "name",
			"--prompt", "prompt",
			"--timezone", "timezone",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cron: cron\n" +
			"enabled: true\n" +
			"features:\n" +
			"  - memory\n" +
			"instructions: instructions\n" +
			"model: model\n" +
			"name: name\n" +
			"prompt: prompt\n" +
			"timezone: timezone\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cells:schedules", "update",
			"--cell-id", "cell_abc123",
			"--schedule-id", "schedule_abc123",
		)
	})
}

func TestCellsSchedulesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:schedules", "list",
			"--cell-id", "cell_abc123",
		)
	})
}

func TestCellsSchedulesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:schedules", "delete",
			"--cell-id", "cell_abc123",
			"--schedule-id", "schedule_abc123",
		)
	})
}

func TestCellsSchedulesTrigger(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cells:schedules", "trigger",
			"--cell-id", "cell_abc123",
			"--schedule-id", "schedule_abc123",
		)
	})
}
