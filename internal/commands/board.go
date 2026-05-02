package commands

import (
	"fmt"

	"github.com/mrfansi/trello-cli/internal/cmdutil"
	"github.com/mrfansi/trello-cli/internal/output"
	"github.com/mrfansi/trello-cli/internal/trello"
	"github.com/spf13/cobra"
)

func boardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "board",
		Short: "Manage boards",
	}
	cmd.AddCommand(boardListCmd(), boardGetCmd(), boardCreateCmd(), boardDeleteCmd())
	return cmd
}

func boardListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List boards for the authenticated user",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetMembersIdBoards(ctx, "me", &trello.GetMembersIdBoardsParams{})
			if err != nil {
				return err
			}
			var boards []map[string]any
			if err := cmdutil.Decode(resp, &boards); err != nil {
				return err
			}
			rows := make([][]string, 0, len(boards))
			for _, b := range boards {
				rows = append(rows, []string{
					strOrEmpty(b["id"]),
					output.Truncate(strOrEmpty(b["name"]), 40),
					strOrEmpty(b["url"]),
				})
			}
			return cx.Renderer.Table([]string{"ID", "Name", "URL"}, rows)
		},
	}
}

func boardGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <board-id>",
		Short: "Get a board by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetBoardsId(ctx, args[0], &trello.GetBoardsIdParams{})
			if err != nil {
				return err
			}
			var b map[string]any
			if err := cmdutil.Decode(resp, &b); err != nil {
				return err
			}
			if jsonOutput {
				return cx.Renderer.Raw(b)
			}
			return cx.Renderer.KeyValue([][2]string{
				{"ID", strOrEmpty(b["id"])},
				{"Name", strOrEmpty(b["name"])},
				{"Description", strOrEmpty(b["desc"])},
				{"URL", strOrEmpty(b["url"])},
				{"Closed", fmt.Sprint(b["closed"])},
				{"Org ID", strOrEmpty(b["idOrganization"])},
			})
		},
	}
}

func boardCreateCmd() *cobra.Command {
	var (
		desc           string
		idOrganization string
		defaultLists   bool
	)
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a board",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			params := &trello.PostBoardsParams{
				Name:         args[0],
				DefaultLists: &defaultLists,
			}
			if desc != "" {
				params.Desc = &desc
			}
			if idOrganization != "" {
				v := trello.TrelloID(idOrganization)
				params.IdOrganization = &v
			}
			resp, err := cx.Client.PostBoards(ctx, params)
			if err != nil {
				return err
			}
			var b map[string]any
			if err := cmdutil.Decode(resp, &b); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "created board %s (%s)\n", strOrEmpty(b["name"]), strOrEmpty(b["id"]))
			return nil
		},
	}
	cmd.Flags().StringVar(&desc, "desc", "", "Board description")
	cmd.Flags().StringVar(&idOrganization, "org", "", "Organization (workspace) ID")
	cmd.Flags().BoolVar(&defaultLists, "default-lists", true, "Create default lists (To Do, Doing, Done)")
	return cmd
}

func boardDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rm <board-id>",
		Short: "Delete a board",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.DeleteBoardsId(ctx, args[0])
			if err != nil {
				return err
			}
			if err := cmdutil.Decode(resp, nil); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "deleted board %s\n", args[0])
			return nil
		},
	}
}

func strOrEmpty(v any) string {
	if v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return fmt.Sprint(v)
	}
	return s
}
