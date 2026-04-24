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

var environmentsWebhooksCreate = cli.Command{
	Name:    "create",
	Usage:   "Perform create operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "HTTPS endpoint that will receive webhook deliveries.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[[]string]{
			Name:     "event",
			Usage:    "Event names to deliver. Omit to subscribe to all non-test events.",
			BodyPath: "events",
		},
	},
	Action:          handleEnvironmentsWebhooksCreate,
	HideHelpCommand: true,
}

var environmentsWebhooksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Perform retrieve operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
	},
	Action:          handleEnvironmentsWebhooksRetrieve,
	HideHelpCommand: true,
}

var environmentsWebhooksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Perform update operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Usage:    "Whether deliveries are enabled for this subscription.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "event",
			Usage:    "Event names to deliver. Omit to subscribe to all non-test events.",
			BodyPath: "events",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "HTTPS endpoint that will receive webhook deliveries.",
			BodyPath: "url",
		},
	},
	Action:          handleEnvironmentsWebhooksUpdate,
	HideHelpCommand: true,
}

var environmentsWebhooksList = cli.Command{
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
	Action:          handleEnvironmentsWebhooksList,
	HideHelpCommand: true,
}

var environmentsWebhooksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Perform delete operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
	},
	Action:          handleEnvironmentsWebhooksDelete,
	HideHelpCommand: true,
}

var environmentsWebhooksRotate = cli.Command{
	Name:    "rotate",
	Usage:   "Perform rotate operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
	},
	Action:          handleEnvironmentsWebhooksRotate,
	HideHelpCommand: true,
}

var environmentsWebhooksTest = cli.Command{
	Name:    "test",
	Usage:   "Perform test operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "webhook-id",
			Required: true,
		},
	},
	Action:          handleEnvironmentsWebhooksTest,
	HideHelpCommand: true,
}

func handleEnvironmentsWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.EnvironmentWebhookNewParams{}

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
	_, err = client.Environments.Webhooks.New(
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
		Title:          "environments:webhooks create",
		Transform:      transform,
	})
}

func handleEnvironmentsWebhooksRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Environments.Webhooks.Get(
		ctx,
		cmd.Value("environment-id").(string),
		cmd.Value("webhook-id").(string),
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
		Title:          "environments:webhooks retrieve",
		Transform:      transform,
	})
}

func handleEnvironmentsWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.EnvironmentWebhookUpdateParams{}

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
	_, err = client.Environments.Webhooks.Update(
		ctx,
		cmd.Value("environment-id").(string),
		cmd.Value("webhook-id").(string),
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
		Title:          "environments:webhooks update",
		Transform:      transform,
	})
}

func handleEnvironmentsWebhooksList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.EnvironmentWebhookListParams{}

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
	_, err = client.Environments.Webhooks.List(
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
		Title:          "environments:webhooks list",
		Transform:      transform,
	})
}

func handleEnvironmentsWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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

	return client.Environments.Webhooks.Delete(
		ctx,
		cmd.Value("environment-id").(string),
		cmd.Value("webhook-id").(string),
		options...,
	)
}

func handleEnvironmentsWebhooksRotate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Environments.Webhooks.Rotate(
		ctx,
		cmd.Value("environment-id").(string),
		cmd.Value("webhook-id").(string),
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
		Title:          "environments:webhooks rotate",
		Transform:      transform,
	})
}

func handleEnvironmentsWebhooksTest(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("environment-id") && len(unusedArgs) > 0 {
		cmd.Set("environment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Environments.Webhooks.Test(
		ctx,
		cmd.Value("environment-id").(string),
		cmd.Value("webhook-id").(string),
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
		Title:          "environments:webhooks test",
		Transform:      transform,
	})
}
