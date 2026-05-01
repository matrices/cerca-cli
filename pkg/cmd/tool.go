// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/matrices/cerca-cli/internal/apiquery"
	"github.com/matrices/cerca-cli/internal/requestflag"
	"github.com/matrices/cerca-go"
	"github.com/matrices/cerca-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var toolsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create tool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fleet-id",
			Required:  true,
			PathParam: "fleetId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "auth",
			Usage:    "Tool source authentication configuration. The `kind` field selects one of `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.",
			Required: true,
			BodyPath: "auth",
		},
		&requestflag.Flag[string]{
			Name:     "namespace",
			Required: true,
			BodyPath: "namespace",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tool",
			Required: true,
			BodyPath: "tools",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "http".`,
			Default:  "http",
			Const:    true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "approval",
			Usage:    `Allowed values: "always", "never".`,
			BodyPath: "approval",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "execution",
			Usage:    "HTTP tool execution retry and timeout policy.",
			BodyPath: "execution",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleToolsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"tool": {
		&requestflag.InnerFlag[string]{
			Name:       "tool.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.endpoint",
			Usage:      "HTTP endpoint invoked when the tool is called.",
			InnerField: "endpoint",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.input-schema",
			Usage:      "JSON Schema object describing tool input parameters.",
			InnerField: "inputSchema",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tool.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tool.approval",
			Usage:      `Allowed values: "always", "never".`,
			InnerField: "approval",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.execution",
			Usage:      "HTTP tool execution retry and timeout policy.",
			InnerField: "execution",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.response",
			Usage:      "How the HTTP response should be normalized for the agent.",
			InnerField: "response",
		},
	},
	"execution": {
		&requestflag.InnerFlag[string]{
			Name:       "execution.idempotency-key-header",
			InnerField: "idempotencyKeyHeader",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "execution.max-attempts",
			InnerField: "maxAttempts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "execution.retry-mode",
			Usage:      `Allowed values: "disabled", "safe_only", "enabled".`,
			InnerField: "retryMode",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "execution.timeout-ms",
			InnerField: "timeoutMs",
		},
	},
})

var toolsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve tool source",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fleet-id",
			Required:  true,
			PathParam: "fleetId",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Required:  true,
			PathParam: "sourceId",
		},
	},
	Action:          handleToolsRetrieve,
	HideHelpCommand: true,
}

var toolsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update tool source",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fleet-id",
			Required:  true,
			PathParam: "fleetId",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Required:  true,
			PathParam: "sourceId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "auth",
			Usage:    "Tool source authentication configuration. The `kind` field selects one of `none`, `api_key`, `oauth_exchange`, or `oauth_connection`.",
			Required: true,
			BodyPath: "auth",
		},
		&requestflag.Flag[string]{
			Name:     "namespace",
			Required: true,
			BodyPath: "namespace",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tool",
			Required: true,
			BodyPath: "tools",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "http".`,
			Default:  "http",
			Const:    true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "approval",
			Usage:    `Allowed values: "always", "never".`,
			BodyPath: "approval",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "execution",
			Usage:    "HTTP tool execution retry and timeout policy.",
			BodyPath: "execution",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleToolsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"tool": {
		&requestflag.InnerFlag[string]{
			Name:       "tool.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.endpoint",
			Usage:      "HTTP endpoint invoked when the tool is called.",
			InnerField: "endpoint",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.input-schema",
			Usage:      "JSON Schema object describing tool input parameters.",
			InnerField: "inputSchema",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tool.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tool.approval",
			Usage:      `Allowed values: "always", "never".`,
			InnerField: "approval",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.execution",
			Usage:      "HTTP tool execution retry and timeout policy.",
			InnerField: "execution",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tool.response",
			Usage:      "How the HTTP response should be normalized for the agent.",
			InnerField: "response",
		},
	},
	"execution": {
		&requestflag.InnerFlag[string]{
			Name:       "execution.idempotency-key-header",
			InnerField: "idempotencyKeyHeader",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "execution.max-attempts",
			InnerField: "maxAttempts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "execution.retry-mode",
			Usage:      `Allowed values: "disabled", "safe_only", "enabled".`,
			InnerField: "retryMode",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "execution.timeout-ms",
			InnerField: "timeoutMs",
		},
	},
})

var toolsList = cli.Command{
	Name:    "list",
	Usage:   "List tools",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fleet-id",
			Required:  true,
			PathParam: "fleetId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor returned by a previous request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			Usage:     "Maximum number of items to return. Defaults to 20 and preserves parseInt semantics.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleToolsList,
	HideHelpCommand: true,
}

var toolsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete tool source",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "fleet-id",
			Required:  true,
			PathParam: "fleetId",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Required:  true,
			PathParam: "sourceId",
		},
	},
	Action:          handleToolsDelete,
	HideHelpCommand: true,
}

func handleToolsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fleet-id") && len(unusedArgs) > 0 {
		cmd.Set("fleet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cercago.ToolNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tools.New(
		ctx,
		cmd.Value("fleet-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tools create",
		Transform:      transform,
	})
}

func handleToolsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fleet-id") && len(unusedArgs) > 0 {
		cmd.Set("fleet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tools.Get(
		ctx,
		cmd.Value("fleet-id").(string),
		cmd.Value("source-id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tools retrieve",
		Transform:      transform,
	})
}

func handleToolsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fleet-id") && len(unusedArgs) > 0 {
		cmd.Set("fleet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cercago.ToolUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tools.Update(
		ctx,
		cmd.Value("fleet-id").(string),
		cmd.Value("source-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tools update",
		Transform:      transform,
	})
}

func handleToolsList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fleet-id") && len(unusedArgs) > 0 {
		cmd.Set("fleet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cercago.ToolListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Tools.List(
			ctx,
			cmd.Value("fleet-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tools list",
			Transform:      transform,
		})
	} else {
		iter := client.Tools.ListAutoPaging(
			ctx,
			cmd.Value("fleet-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tools list",
			Transform:      transform,
		})
	}
}

func handleToolsDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fleet-id") && len(unusedArgs) > 0 {
		cmd.Set("fleet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Tools.Delete(
		ctx,
		cmd.Value("fleet-id").(string),
		cmd.Value("source-id").(string),
		options...,
	)
}
