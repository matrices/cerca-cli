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

var sandboxExec = cli.Command{
	Name:    "exec",
	Usage:   "Perform exec operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "command",
			Required: true,
			BodyPath: "command",
		},
		&requestflag.Flag[float64]{
			Name:     "max-buffer",
			BodyPath: "maxBuffer",
		},
		&requestflag.Flag[float64]{
			Name:     "execution-timeout",
			Usage:    "Timeout in seconds. Runtime converts this to milliseconds.",
			BodyPath: "timeout",
		},
		&requestflag.Flag[*string]{
			Name:     "workdir",
			Usage:    "Optional sandbox working directory.",
			BodyPath: "workdir",
		},
	},
	Action:          handleSandboxExec,
	HideHelpCommand: true,
}

var sandboxRead = cli.Command{
	Name:    "read",
	Usage:   "Perform read operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "path",
			Required: true,
			BodyPath: "path",
		},
		&requestflag.Flag[float64]{
			Name:     "limit",
			BodyPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:     "offset",
			BodyPath: "offset",
		},
	},
	Action:          handleSandboxRead,
	HideHelpCommand: true,
}

var sandboxWrite = cli.Command{
	Name:    "write",
	Usage:   "Perform write operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "path",
			Required: true,
			BodyPath: "path",
		},
	},
	Action:          handleSandboxWrite,
	HideHelpCommand: true,
}

func handleSandboxExec(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.SandboxExecParams{}

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
	_, err = client.Sandbox.Exec(
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
		Title:          "sandbox exec",
		Transform:      transform,
	})
}

func handleSandboxRead(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.SandboxReadParams{}

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
	_, err = client.Sandbox.Read(
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
		Title:          "sandbox read",
		Transform:      transform,
	})
}

func handleSandboxWrite(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cercago.SandboxWriteParams{}

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
	_, err = client.Sandbox.Write(
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
		Title:          "sandbox write",
		Transform:      transform,
	})
}
