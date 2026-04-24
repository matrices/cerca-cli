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

var cellsThreadsEventsList = cli.Command{
	Name:    "list",
	Usage:   "Perform list operation",
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
			Name:      "cursor",
			Usage:     "Opaque pagination cursor returned by a previous request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "events",
			Usage:     "Comma-separated event type filter.",
			QueryPath: "events",
		},
		&requestflag.Flag[string]{
			Name:      "history",
			Usage:     "When true, starts from the beginning of the retained buffer.",
			QueryPath: "history",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			Usage:     "Maximum number of items to return. Defaults to 20 and preserves parseInt semantics.",
			QueryPath: "limit",
		},
	},
	Action:          handleCellsThreadsEventsList,
	HideHelpCommand: true,
}

var cellsThreadsEventsSubscribe = cli.Command{
	Name:    "subscribe",
	Usage:   "WebSocket upgrade endpoint. Set\n`Sec-WebSocket-Protocol: cell-v1, cell-auth-<API_KEY>` so the runtime can\nauthenticate the stream while preserving the public subprotocol. HTTP clients\nthat cannot upgrade should use `/cells/{cellId}/threads/{threadId}/events` as\nthe polling analog.",
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
	},
	Action:          handleCellsThreadsEventsSubscribe,
	HideHelpCommand: true,
}

func handleCellsThreadsEventsList(ctx context.Context, cmd *cli.Command) error {
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
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.CellThreadEventListParams{}

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
	_, err = client.Cells.Threads.Events.List(
		ctx,
		cmd.Value("cell-id").(string),
		cmd.Value("thread-id").(string),
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
		Title:          "cells:threads:events list",
		Transform:      transform,
	})
}

func handleCellsThreadsEventsSubscribe(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Cells.Threads.Events.Subscribe(
		ctx,
		cmd.Value("cell-id").(string),
		cmd.Value("thread-id").(string),
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
		Title:          "cells:threads:events subscribe",
		Transform:      transform,
	})
}
