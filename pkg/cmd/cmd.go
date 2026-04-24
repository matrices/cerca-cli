// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/cerca-cli/internal/autocomplete"
	"github.com/stainless-sdks/cerca-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "cerca",
		Usage:     "CLI for the cerca API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("CERCA_API_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "auth",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&authContext,
					&authListEnvironments,
				},
			},
			{
				Name:     "oauth",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&oauthConnect,
				},
			},
			{
				Name:     "credentials",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&credentialsList,
					&credentialsDelete,
					&credentialsCreateAPIKey,
				},
			},
			{
				Name:     "environments:tool-sources",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&environmentsToolSourcesCreate,
					&environmentsToolSourcesRetrieve,
					&environmentsToolSourcesUpdate,
					&environmentsToolSourcesList,
					&environmentsToolSourcesDelete,
				},
			},
			{
				Name:     "environments:webhooks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&environmentsWebhooksCreate,
					&environmentsWebhooksRetrieve,
					&environmentsWebhooksUpdate,
					&environmentsWebhooksList,
					&environmentsWebhooksDelete,
					&environmentsWebhooksRotate,
					&environmentsWebhooksTest,
				},
			},
			{
				Name:     "environments:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&environmentsEventsList,
					&environmentsEventsSubscribe,
				},
			},
			{
				Name:     "cells",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsCreate,
					&cellsRetrieve,
					&cellsUpdate,
					&cellsList,
					&cellsDelete,
				},
			},
			{
				Name:     "cells:config",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsConfigRetrieve,
				},
			},
			{
				Name:     "cells:metadata",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsMetadataUpdate,
				},
			},
			{
				Name:     "cells:tools",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsToolsList,
				},
			},
			{
				Name:     "cells:tool-sources",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsToolSourcesList,
				},
			},
			{
				Name:     "cells:context",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsContextList,
					&cellsContextSearch,
				},
			},
			{
				Name:     "cells:context:entries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsContextEntriesRetrieve,
					&cellsContextEntriesDelete,
					&cellsContextEntriesPut,
				},
			},
			{
				Name:     "cells:connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsConnectionsList,
					&cellsConnectionsAttach,
					&cellsConnectionsDetach,
				},
			},
			{
				Name:     "cells:schedules",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsSchedulesCreate,
					&cellsSchedulesUpdate,
					&cellsSchedulesList,
					&cellsSchedulesDelete,
					&cellsSchedulesTrigger,
				},
			},
			{
				Name:     "cells:threads",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsThreadsCreate,
					&cellsThreadsRetrieve,
					&cellsThreadsList,
					&cellsThreadsCancel,
					&cellsThreadsClose,
					&cellsThreadsCompact,
					&cellsThreadsSteer,
					&cellsThreadsTurn,
				},
			},
			{
				Name:     "cells:threads:approvals",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsThreadsApprovalsResolve,
				},
			},
			{
				Name:     "cells:threads:approval-grants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsThreadsApprovalGrantsList,
					&cellsThreadsApprovalGrantsDelete,
				},
			},
			{
				Name:     "cells:threads:logs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsThreadsLogsList,
				},
			},
			{
				Name:     "cells:threads:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsThreadsEventsList,
					&cellsThreadsEventsSubscribe,
				},
			},
			{
				Name:     "cells:approvals",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsApprovalsList,
				},
			},
			{
				Name:     "cells:approval-grants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsApprovalGrantsList,
					&cellsApprovalGrantsDelete,
				},
			},
			{
				Name:     "cells:logs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsLogsList,
				},
			},
			{
				Name:     "cells:sandbox",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsSandboxExec,
					&cellsSandboxRead,
					&cellsSandboxWrite,
				},
			},
			{
				Name:     "cells:events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cellsEventsList,
					&cellsEventsSubscribe,
				},
			},
			{
				Name:     "models",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&modelsList,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "cerca @manpages [-o cerca.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "cerca.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "cerca.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
