// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/matrices/cerca-go"
	"github.com/matrices/cerca-go/option"
	"github.com/stainless-sdks/cerca-cli/internal/apiquery"
	"github.com/stainless-sdks/cerca-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var environmentsToolSourcesCreate = cli.Command{
	Name:    "create",
	Usage:   "Perform create operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[any]{
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
		&requestflag.Flag[[]any]{
			Name:     "tool",
			Required: true,
			BodyPath: "tools",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "http".`,
			Required: true,
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
		&requestflag.Flag[any]{
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
	Action:          handleEnvironmentsToolSourcesCreate,
	HideHelpCommand: true,
}

var environmentsToolSourcesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Perform retrieve operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "source-id",
			Required: true,
		},
	},
	Action:          handleEnvironmentsToolSourcesRetrieve,
	HideHelpCommand: true,
}

var environmentsToolSourcesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Perform update operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "source-id",
			Required: true,
		},
		&requestflag.Flag[any]{
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
		&requestflag.Flag[[]any]{
			Name:     "tool",
			Required: true,
			BodyPath: "tools",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "http".`,
			Required: true,
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
		&requestflag.Flag[any]{
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
	Action:          handleEnvironmentsToolSourcesUpdate,
	HideHelpCommand: true,
}

var environmentsToolSourcesList = cli.Command{
	Name:    "list",
	Usage:   "Perform list operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
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
	},
	Action:          handleEnvironmentsToolSourcesList,
	HideHelpCommand: true,
}

var environmentsToolSourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Perform delete operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "source-id",
			Required: true,
		},
	},
	Action:          handleEnvironmentsToolSourcesDelete,
	HideHelpCommand: true,
}

func handleEnvironmentsToolSourcesCreate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.EnvironmentToolSourceNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Environments.ToolSources.New(
		ctx,
		cmd.Value("environment-id").(string),
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
		Title:          "environments:tool-sources create",
		Transform:      transform,
	})
}

func handleEnvironmentsToolSourcesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
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
	_, err = client.Environments.ToolSources.Get(
		ctx,
		cmd.Value("environment-id").(string),
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
		Title:          "environments:tool-sources retrieve",
		Transform:      transform,
	})
}

func handleEnvironmentsToolSourcesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.EnvironmentToolSourceUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Environments.ToolSources.Update(
		ctx,
		cmd.Value("environment-id").(string),
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
		Title:          "environments:tool-sources update",
		Transform:      transform,
	})
}

func handleEnvironmentsToolSourcesList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.EnvironmentToolSourceListParams{}

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
	_, err = client.Environments.ToolSources.List(
		ctx,
		cmd.Value("environment-id").(string),
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
		Title:          "environments:tool-sources list",
		Transform:      transform,
	})
}

func handleEnvironmentsToolSourcesDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
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

	return client.Environments.ToolSources.Delete(
		ctx,
		cmd.Value("environment-id").(string),
		cmd.Value("source-id").(string),
		options...,
	)
}
