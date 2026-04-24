// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cerca-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "create",
			"--fleet-id", "fleet_abc123",
			"--url", "https://example.com/webhook",
			"--event", "agent.created",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com/webhook\n" +
			"events:\n" +
			"  - agent.created\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks", "create",
			"--fleet-id", "fleet_abc123",
		)
	})
}

func TestWebhooksRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "retrieve",
			"--fleet-id", "fleet_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestWebhooksUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "update",
			"--fleet-id", "fleet_abc123",
			"--webhook-id", "webhook_abc123",
			"--enabled=true",
			"--event", "agent.created",
			"--url", "https://example.com/webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"events:\n" +
			"  - agent.created\n" +
			"url: https://example.com/webhook\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks", "update",
			"--fleet-id", "fleet_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestWebhooksList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "list",
			"--max-items", "10",
			"--fleet-id", "fleet_abc123",
			"--cursor", "cursor_abc123",
			"--limit", "20",
		)
	})
}

func TestWebhooksDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "delete",
			"--fleet-id", "fleet_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestWebhooksRotate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "rotate",
			"--fleet-id", "fleet_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}

func TestWebhooksTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "test",
			"--fleet-id", "fleet_abc123",
			"--webhook-id", "webhook_abc123",
		)
	})
}
