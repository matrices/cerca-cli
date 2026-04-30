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

var schedulesCreate = cli.Command{
	Name:    "create",
	Usage:   "Schedules",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "cron",
			BodyPath: "cron",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "run-at",
			BodyPath: "runAt",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			BodyPath: "timezone",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "Per-schedule tool subset. When the schedule starts a thread, these tools can only narrow the agent's effective tools.",
			BodyPath: "tools",
		},
	},
	Action:          handleSchedulesCreate,
	HideHelpCommand: true,
}

var schedulesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Schedule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "schedule-id",
			Required:  true,
			PathParam: "scheduleId",
		},
		&requestflag.Flag[string]{
			Name:     "cron",
			BodyPath: "cron",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
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
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "run-at",
			BodyPath: "runAt",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			BodyPath: "timezone",
		},
		&requestflag.Flag[[]string]{
			Name:     "tool",
			Usage:    "Per-schedule tool subset. When updated, these tools can only narrow the agent's effective tools for future scheduled threads.",
			BodyPath: "tools",
		},
	},
	Action:          handleSchedulesUpdate,
	HideHelpCommand: true,
}

var schedulesList = cli.Command{
	Name:    "list",
	Usage:   "Schedules",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
	},
	Action:          handleSchedulesList,
	HideHelpCommand: true,
}

var schedulesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Schedule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "schedule-id",
			Required:  true,
			PathParam: "scheduleId",
		},
	},
	Action:          handleSchedulesDelete,
	HideHelpCommand: true,
}

var schedulesTrigger = cli.Command{
	Name:    "trigger",
	Usage:   "Trigger",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "schedule-id",
			Required:  true,
			PathParam: "scheduleId",
		},
	},
	Action:          handleSchedulesTrigger,
	HideHelpCommand: true,
}

func handleSchedulesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := cercago.ScheduleNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Schedules.New(
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
		Title:          "schedules create",
		Transform:      transform,
	})
}

func handleSchedulesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("schedule-id") && len(unusedArgs) > 0 {
		cmd.Set("schedule-id", unusedArgs[0])
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

	params := cercago.ScheduleUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Schedules.Update(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("schedule-id").(string),
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
		Title:          "schedules update",
		Transform:      transform,
	})
}

func handleSchedulesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Schedules.List(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "schedules list",
		Transform:      transform,
	})
}

func handleSchedulesDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("schedule-id") && len(unusedArgs) > 0 {
		cmd.Set("schedule-id", unusedArgs[0])
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
	_, err = client.Schedules.Delete(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("schedule-id").(string),
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
		Title:          "schedules delete",
		Transform:      transform,
	})
}

func handleSchedulesTrigger(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("schedule-id") && len(unusedArgs) > 0 {
		cmd.Set("schedule-id", unusedArgs[0])
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
	_, err = client.Schedules.Trigger(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("schedule-id").(string),
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
		Title:          "schedules trigger",
		Transform:      transform,
	})
}
