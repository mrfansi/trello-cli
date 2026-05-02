package commands

import (
	"fmt"

	"github.com/mrfansi/trello-cli/internal/cmdutil"
	"github.com/mrfansi/trello-cli/internal/trello"
	"github.com/spf13/cobra"
)

func memberCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member",
		Short: "Look up members",
	}
	cmd.AddCommand(memberGetCmd())
	return cmd
}

func memberGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <id-or-username>",
		Short: "Get a member by id or username",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetMembersId(ctx, args[0], &trello.GetMembersIdParams{})
			if err != nil {
				return err
			}
			var m map[string]any
			if err := cmdutil.Decode(resp, &m); err != nil {
				return err
			}
			if jsonOutput {
				return cx.Renderer.Raw(m)
			}
			return cx.Renderer.KeyValue([][2]string{
				{"ID", strOrEmpty(m["id"])},
				{"Username", strOrEmpty(m["username"])},
				{"Full Name", strOrEmpty(m["fullName"])},
				{"Email", strOrEmpty(m["email"])},
				{"URL", strOrEmpty(m["url"])},
				{"Bio", strOrEmpty(m["bio"])},
				{"Confirmed", fmt.Sprint(m["confirmed"])},
			})
		},
	}
}

func meCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "me",
		Short: "Show authenticated user",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetMembersId(ctx, "me", &trello.GetMembersIdParams{})
			if err != nil {
				return err
			}
			var m map[string]any
			if err := cmdutil.Decode(resp, &m); err != nil {
				return err
			}
			if jsonOutput {
				return cx.Renderer.Raw(m)
			}
			return cx.Renderer.KeyValue([][2]string{
				{"ID", strOrEmpty(m["id"])},
				{"Username", strOrEmpty(m["username"])},
				{"Full Name", strOrEmpty(m["fullName"])},
				{"Email", strOrEmpty(m["email"])},
				{"URL", strOrEmpty(m["url"])},
			})
		},
	}
}
