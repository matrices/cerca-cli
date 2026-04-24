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

var cellsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Perform create operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
			BodyPath: "userId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "configuration",
			BodyPath: "configuration",
		},
		&requestflag.Flag[string]{
			Name:     "environment-id",
			BodyPath: "environmentId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary string metadata stored on a cell. Runtime enforces maximum key and value sizes.",
			BodyPath: "metadata",
		},
	},
	Action:          handleCellsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.approval-required",
			InnerField: "approvalRequired",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "configuration.approval-timeout-ms",
			InnerField: "approvalTimeoutMs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.default-model",
			InnerField: "defaultModel",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.excluded-tool-source-ids",
			InnerField: "excludedToolSourceIds",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.features",
			InnerField: "features",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.instructions",
			InnerField: "instructions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.tool-approval-overrides",
			InnerField: "toolApprovalOverrides",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.tool-source-mode",
			Usage:      `Allowed values: "inherit", "explicit".`,
			InnerField: "toolSourceMode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.user-context",
			InnerField: "userContext",
		},
	},
})

var cellsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Perform retrieve operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
	},
	Action:          handleCellsRetrieve,
	HideHelpCommand: true,
}

var cellsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Perform update operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "configuration",
			Required: true,
			BodyPath: "configuration",
		},
	},
	Action:          handleCellsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.approval-required",
			InnerField: "approvalRequired",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "configuration.approval-timeout-ms",
			InnerField: "approvalTimeoutMs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.default-model",
			InnerField: "defaultModel",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.excluded-tool-source-ids",
			InnerField: "excludedToolSourceIds",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.features",
			InnerField: "features",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.instructions",
			InnerField: "instructions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.tool-approval-overrides",
			InnerField: "toolApprovalOverrides",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.tool-source-mode",
			Usage:      `Allowed values: "inherit", "explicit".`,
			InnerField: "toolSourceMode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.user-context",
			InnerField: "userContext",
		},
	},
})

var cellsList = cli.Command{
	Name:    "list",
	Usage:   "Perform list operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "active",
			Usage:     "When set to true, lists only cells with active or awaiting threads.",
			QueryPath: "active",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor returned by a previous request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "environment-id",
			Usage:     "Environment to list cells from. Defaults to the API key's default environment when omitted.",
			QueryPath: "environmentId",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			Usage:     "Maximum number of items to return. Defaults to 20 and preserves parseInt semantics.",
			QueryPath: "limit",
		},
	},
	Action:          handleCellsList,
	HideHelpCommand: true,
}

var cellsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Perform delete operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
	},
	Action:          handleCellsDelete,
	HideHelpCommand: true,
}

func handleCellsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellNewParams{}

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
	_, err = client.Cells.New(ctx, params, options...)
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
		Title:          "cells create",
		Transform:      transform,
	})
}

func handleCellsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
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
	_, err = client.Cells.Get(ctx, cmd.Value("cell-id").(string), options...)
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
		Title:          "cells retrieve",
		Transform:      transform,
	})
}

func handleCellsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellUpdateParams{}

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
	_, err = client.Cells.Update(
		ctx,
		cmd.Value("cell-id").(string),
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
		Title:          "cells update",
		Transform:      transform,
	})
}

func handleCellsList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellListParams{}

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
	_, err = client.Cells.List(ctx, params, options...)
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
		Title:          "cells list",
		Transform:      transform,
	})
}

func handleCellsDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
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
	_, err = client.Cells.Delete(ctx, cmd.Value("cell-id").(string), options...)
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
		Title:          "cells delete",
		Transform:      transform,
	})
}
