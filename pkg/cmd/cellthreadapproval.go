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

var cellsThreadsApprovalsResolve = cli.Command{
	Name:    "resolve",
	Usage:   "Perform resolve operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "thread-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "approval-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "decision",
			Usage:    `Allowed values: "approve", "deny", "cancel".`,
			Required: true,
			BodyPath: "decision",
		},
		&requestflag.Flag[string]{
			Name:     "grant",
			Usage:    `Allowed values: "thread", "cell".`,
			BodyPath: "grant",
		},
	},
	Action:          handleCellsThreadsApprovalsResolve,
	HideHelpCommand: true,
}

func handleCellsThreadsApprovalsResolve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("thread-id") && len(unusedArgs) > 0 {
		cmd.Set("thread-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("approval-id") && len(unusedArgs) > 0 {
		cmd.Set("approval-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellThreadApprovalResolveParams{}

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
	_, err = client.Cells.Threads.Approvals.Resolve(
		ctx,
		cmd.Value("cell-id").(string),
		cmd.Value("thread-id").(string),
		cmd.Value("approval-id").(string),
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
		Title:          "cells:threads:approvals resolve",
		Transform:      transform,
	})
}
