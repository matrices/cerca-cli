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

var cellsApprovalGrantsList = cli.Command{
	Name:    "list",
	Usage:   "Perform list operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
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
	Action:          handleCellsApprovalGrantsList,
	HideHelpCommand: true,
}

var cellsApprovalGrantsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Perform delete operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "grant-id",
			Required: true,
		},
	},
	Action:          handleCellsApprovalGrantsDelete,
	HideHelpCommand: true,
}

func handleCellsApprovalGrantsList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellApprovalGrantListParams{}

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
	_, err = client.Cells.ApprovalGrants.List(
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
		Title:          "cells:approval-grants list",
		Transform:      transform,
	})
}

func handleCellsApprovalGrantsDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("grant-id") && len(unusedArgs) > 0 {
		cmd.Set("grant-id", unusedArgs[0])
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
	_, err = client.Cells.ApprovalGrants.Delete(
		ctx,
		cmd.Value("cell-id").(string),
		cmd.Value("grant-id").(string),
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
		Title:          "cells:approval-grants delete",
		Transform:      transform,
	})
}
