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

var agentsCreate = requestflag.WithInnerFlags(cli.Command{
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
			Name:     "fleet-id",
			BodyPath: "fleetId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary string metadata stored on an agent. Runtime enforces maximum key and value sizes.",
			BodyPath: "metadata",
		},
	},
	Action:          handleAgentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.approvals",
			InnerField: "approvals",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.default-model",
			InnerField: "defaultModel",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.instructions",
			InnerField: "instructions",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.tools",
			InnerField: "tools",
		},
	},
})

var agentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Perform retrieve operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
	},
	Action:          handleAgentsRetrieve,
	HideHelpCommand: true,
}

var agentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Perform update operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "configuration",
			Required: true,
			BodyPath: "configuration",
		},
	},
	Action:          handleAgentsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.approvals",
			InnerField: "approvals",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.default-model",
			InnerField: "defaultModel",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.instructions",
			InnerField: "instructions",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "configuration.tools",
			InnerField: "tools",
		},
	},
})

var agentsList = cli.Command{
	Name:    "list",
	Usage:   "Perform list operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "active",
			Usage:     "When set to true, lists only agents with active or awaiting threads.",
			QueryPath: "active",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor returned by a previous request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "fleet-id",
			Usage:     "Fleet to list agents from. Defaults to the API key's default fleet when omitted.",
			QueryPath: "fleetId",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			Usage:     "Maximum number of items to return. Defaults to 20 and preserves parseInt semantics.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentsList,
	HideHelpCommand: true,
}

var agentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Perform delete operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
	},
	Action:          handleAgentsDelete,
	HideHelpCommand: true,
}

var agentsRetrieveConfig = cli.Command{
	Name:    "retrieve-config",
	Usage:   "Perform retrieve-config operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
	},
	Action:          handleAgentsRetrieveConfig,
	HideHelpCommand: true,
}

var agentsUpdateMetadata = cli.Command{
	Name:    "update-metadata",
	Usage:   "Perform update-metadata operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Arbitrary string metadata stored on an agent. Runtime enforces maximum key and value sizes.",
			Required: true,
			BodyPath: "metadata",
		},
	},
	Action:          handleAgentsUpdateMetadata,
	HideHelpCommand: true,
}

func handleAgentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.AgentNewParams{}

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
	_, err = client.Agents.New(ctx, params, options...)
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
		Title:          "agents create",
		Transform:      transform,
	})
}

func handleAgentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.Get(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "agents retrieve",
		Transform:      transform,
	})
}

func handleAgentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.AgentUpdateParams{}

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
	_, err = client.Agents.Update(
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
		Title:          "agents update",
		Transform:      transform,
	})
}

func handleAgentsList(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.AgentListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Agents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "agents list",
			Transform:      transform,
		})
	} else {
		iter := client.Agents.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "agents list",
			Transform:      transform,
		})
	}
}

func handleAgentsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.Delete(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "agents delete",
		Transform:      transform,
	})
}

func handleAgentsRetrieveConfig(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.GetConfig(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "agents retrieve-config",
		Transform:      transform,
	})
}

func handleAgentsUpdateMetadata(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.AgentUpdateMetadataParams{}

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
	_, err = client.Agents.UpdateMetadata(
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
		Title:          "agents update-metadata",
		Transform:      transform,
	})
}
