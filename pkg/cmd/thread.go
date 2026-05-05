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

var threadsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create thread",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "message",
			BodyPath: "message",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "system-prompt",
			Usage:    "Deprecated alias for `instructions`; accepted for backwards compatibility.",
			BodyPath: "systemPrompt",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "Per-thread tool subset. Omit to inherit the agent's full effective tools; pass [] to run with no configurable tools. Provided entries can only narrow the agent's effective tools.",
			BodyPath: "tools",
		},
	},
	Action:          handleThreadsCreate,
	HideHelpCommand: true,
}

var threadsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve thread",
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
			Name:      "debug",
			Usage:     "When true, includes debug-only compiled context fields.",
			QueryPath: "debug",
		},
		&requestflag.Flag[string]{
			Name:      "include-messages",
			Usage:     "When true, includes message content in the thread detail.",
			QueryPath: "includeMessages",
		},
	},
	Action:          handleThreadsRetrieve,
	HideHelpCommand: true,
}

var threadsList = cli.Command{
	Name:    "list",
	Usage:   "List threads",
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
			Name:      "schedule-id",
			Usage:     "Optional schedule id filter.",
			QueryPath: "scheduleId",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "`idle` threads can accept a new turn or be closed. `running` threads have an active turn. `awaiting` threads are paused on external input such as approvals. `closed` threads are terminal.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleThreadsList,
	HideHelpCommand: true,
}

var threadsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a running or awaiting thread. The underlying runtime treats repeat\ncancellation as an idempotent lifecycle operation when possible.",
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
	},
	Action:          handleThreadsCancel,
	HideHelpCommand: true,
}

var threadsClose = cli.Command{
	Name:    "close",
	Usage:   "Close an idle thread. Closing a running, awaiting, or already-closed thread\nreturns a lifecycle conflict.",
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
	},
	Action:          handleThreadsClose,
	HideHelpCommand: true,
}

var threadsCompact = cli.Command{
	Name:    "compact",
	Usage:   "Force context compaction for an idle thread. Compacting a running thread returns\na lifecycle conflict.",
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
	},
	Action:          handleThreadsCompact,
	HideHelpCommand: true,
}

var threadsStartTurn = cli.Command{
	Name:    "start-turn",
	Usage:   "Create turn",
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
			Name:     "message",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "Per-turn tool subset. Omit to inherit the thread's current tools; pass [] to run this turn with no configurable tools. Provided entries can only narrow the agent/thread effective tools.",
			BodyPath: "tools",
		},
	},
	Action:          handleThreadsStartTurn,
	HideHelpCommand: true,
}

var threadsSteer = cli.Command{
	Name:    "steer",
	Usage:   "Steer a thread with another user message. Steering a closed thread returns a\nconflict; steering a running or awaiting thread queues the message.",
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
			Name:     "message",
			Required: true,
			BodyPath: "message",
		},
	},
	Action:          handleThreadsSteer,
	HideHelpCommand: true,
}

func handleThreadsCreate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cercago.ThreadNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Threads.New(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads create",
		Transform:      transform,
	})
}

func handleThreadsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := cercago.ThreadGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Threads.Get(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads retrieve",
		Transform:      transform,
	})
}

func handleThreadsList(ctx context.Context, cmd *cli.Command) error {
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

	params := cercago.ThreadListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Threads.List(
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
			Title:          "threads list",
			Transform:      transform,
		})
	} else {
		iter := client.Threads.ListAutoPaging(
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
			Title:          "threads list",
			Transform:      transform,
		})
	}
}

func handleThreadsCancel(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Threads.Cancel(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads cancel",
		Transform:      transform,
	})
}

func handleThreadsClose(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Threads.Close(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads close",
		Transform:      transform,
	})
}

func handleThreadsCompact(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Threads.Compact(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads compact",
		Transform:      transform,
	})
}

func handleThreadsStartTurn(ctx context.Context, cmd *cli.Command) error {
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

	params := cercago.ThreadStartTurnParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Threads.StartTurn(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads start-turn",
		Transform:      transform,
	})
}

func handleThreadsSteer(ctx context.Context, cmd *cli.Command) error {
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

	params := cercago.ThreadSteerParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Threads.Steer(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "threads steer",
		Transform:      transform,
	})
}
