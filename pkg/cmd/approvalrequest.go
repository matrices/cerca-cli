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

var approvalRequestsList = cli.Command{
	Name:    "list",
	Usage:   "List approvals",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
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
		&requestflag.Flag[string]{
			Name:      "thread-id",
			Usage:     "Optional thread id filter.",
			QueryPath: "threadId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleApprovalRequestsList,
	HideHelpCommand: true,
}

var approvalRequestsResolve = cli.Command{
	Name:    "resolve",
	Usage:   "Create approval",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "thread-id",
			Required:  true,
			PathParam: "threadId",
		},
		&requestflag.Flag[string]{
			Name:      "approval-id",
			Required:  true,
			PathParam: "approvalId",
		},
		&requestflag.Flag[string]{
			Name:     "decision",
			Usage:    `Allowed values: "approve", "deny", "cancel".`,
			Required: true,
			BodyPath: "decision",
		},
		&requestflag.Flag[string]{
			Name:     "grant",
			Usage:    `Allowed values: "thread", "agent".`,
			BodyPath: "grant",
		},
	},
	Action:          handleApprovalRequestsResolve,
	HideHelpCommand: true,
}

func handleApprovalRequestsList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cercago.ApprovalRequestListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.ApprovalRequests.List(
			ctx,
			cmd.Value("agent-id").(string),
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
			Title:          "approval-requests list",
			Transform:      transform,
		})
	} else {
		iter := client.ApprovalRequests.ListAutoPaging(
			ctx,
			cmd.Value("agent-id").(string),
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
			Title:          "approval-requests list",
			Transform:      transform,
		})
	}
}

func handleApprovalRequestsResolve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cercago.ApprovalRequestResolveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ApprovalRequests.Resolve(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "approval-requests resolve",
		Transform:      transform,
	})
}
