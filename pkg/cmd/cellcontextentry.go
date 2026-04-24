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

var cellsContextEntriesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Perform retrieve operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Usage:     "Context entry key.",
			Required:  true,
			QueryPath: "key",
		},
	},
	Action:          handleCellsContextEntriesRetrieve,
	HideHelpCommand: true,
}

var cellsContextEntriesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Perform delete operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "key",
			Usage:     "Context entry key.",
			Required:  true,
			QueryPath: "key",
		},
	},
	Action:          handleCellsContextEntriesDelete,
	HideHelpCommand: true,
}

var cellsContextEntriesPut = cli.Command{
	Name:    "put",
	Usage:   "Perform put operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cell-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "key",
			Required: true,
			BodyPath: "key",
		},
		&requestflag.Flag[float64]{
			Name:     "expected-version",
			BodyPath: "expectedVersion",
		},
		&requestflag.Flag[string]{
			Name:     "mime-type",
			BodyPath: "mimeType",
		},
	},
	Action:          handleCellsContextEntriesPut,
	HideHelpCommand: true,
}

func handleCellsContextEntriesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellContextEntryGetParams{}

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
	_, err = client.Cells.Context.Entries.Get(
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
		Title:          "cells:context:entries retrieve",
		Transform:      transform,
	})
}

func handleCellsContextEntriesDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellContextEntryDeleteParams{}

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
	_, err = client.Cells.Context.Entries.Delete(
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
		Title:          "cells:context:entries delete",
		Transform:      transform,
	})
}

func handleCellsContextEntriesPut(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cell-id") && len(unusedArgs) > 0 {
		cmd.Set("cell-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellContextEntryPutParams{}

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
	_, err = client.Cells.Context.Entries.Put(
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
		Title:          "cells:context:entries put",
		Transform:      transform,
	})
}
