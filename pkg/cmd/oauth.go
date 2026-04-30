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

var oauthConnect = cli.Command{
	Name:    "connect",
	Usage:   "Perform connect operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "provider",
			Required:  true,
			PathParam: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "return-origin",
			Usage:    "HTTP(S) origin that receives the OAuth completion message.",
			Required: true,
			BodyPath: "returnOrigin",
		},
		&requestflag.Flag[string]{
			Name:     "scope",
			Usage:    "Credential connection scope to attach the OAuth account to.",
			Required: true,
			BodyPath: "scope",
		},
		&requestflag.Flag[[]string]{
			Name:     "scope",
			Usage:    "Provider-specific OAuth scopes. Empty entries are ignored after trimming.",
			BodyPath: "scopes",
		},
	},
	Action:          handleOAuthConnect,
	HideHelpCommand: true,
}

func handleOAuthConnect(ctx context.Context, cmd *cli.Command) error {
	client := cercago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("provider") && len(unusedArgs) > 0 {
		cmd.Set("provider", unusedArgs[0])
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

	params := cercago.OAuthConnectParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.OAuth.Connect(
		ctx,
		cmd.Value("provider").(string),
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
		Title:          "oauth connect",
		Transform:      transform,
	})
}
