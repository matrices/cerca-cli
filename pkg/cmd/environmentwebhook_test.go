// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestEnvironmentsWebhooksCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "create",
			"--environment-id", "env_abc123",
			"--url", "https://example.com/webhook",
			"--event", "cell.created",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com/webhook\n" +
			"events:\n" +
			"  - cell.created\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"environments:webhooks", "create",
			"--environment-id", "env_abc123",
		)
	})
}

func TestEnvironmentsWebhooksRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "retrieve",
			"--environment-id", "env_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestEnvironmentsWebhooksUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "update",
			"--environment-id", "env_abc123",
			"--webhook-id", "webhook_abc123",
			"--enabled=true",
			"--event", "cell.created",
			"--url", "https://example.com/webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"events:\n" +
			"  - cell.created\n" +
			"url: https://example.com/webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"environments:webhooks", "update",
			"--environment-id", "env_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestEnvironmentsWebhooksList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "list",
			"--environment-id", "env_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestEnvironmentsWebhooksDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "delete",
			"--environment-id", "env_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestEnvironmentsWebhooksRotate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "rotate",
			"--environment-id", "env_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestEnvironmentsWebhooksTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"environments:webhooks", "test",
			"--environment-id", "env_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}
