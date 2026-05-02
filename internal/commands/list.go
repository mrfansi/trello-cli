package commands

import (
	"fmt"

	"github.com/mrfansi/trello-cli/internal/cmdutil"
	"github.com/mrfansi/trello-cli/internal/output"
	"github.com/mrfansi/trello-cli/internal/trello"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Manage lists (columns)",
	}
	cmd.AddCommand(listLsCmd(), listGetCmd(), listCreateCmd(), listArchiveCmd())
	return cmd
}

func listLsCmd() *cobra.Command {
	var boardID string
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List lists on a board",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if boardID == "" {
				return fmt.Errorf("--board required")
			}
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetBoardsIdLists(ctx, boardID, &trello.GetBoardsIdListsParams{})
			if err != nil {
				return err
			}
			var lists []map[string]any
			if err := cmdutil.Decode(resp, &lists); err != nil {
				return err
			}
			rows := make([][]string, 0, len(lists))
			for _, l := range lists {
				rows = append(rows, []string{
					strOrEmpty(l["id"]),
					output.Truncate(strOrEmpty(l["name"]), 40),
					fmt.Sprint(l["pos"]),
					fmt.Sprint(l["closed"]),
				})
			}
			return cx.Renderer.Table([]string{"ID", "Name", "Pos", "Closed"}, rows)
		},
	}
	cmd.Flags().StringVarP(&boardID, "board", "b", "", "Board ID (required)")
	return cmd
}

func listGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <list-id>",
		Short: "Get a list by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetListsId(ctx, args[0], &trello.GetListsIdParams{})
			if err != nil {
				return err
			}
			var l map[string]any
			if err := cmdutil.Decode(resp, &l); err != nil {
				return err
			}
			if jsonOutput {
				return cx.Renderer.Raw(l)
			}
			return cx.Renderer.KeyValue([][2]string{
				{"ID", strOrEmpty(l["id"])},
				{"Name", strOrEmpty(l["name"])},
				{"Board ID", strOrEmpty(l["idBoard"])},
				{"Pos", fmt.Sprint(l["pos"])},
				{"Closed", fmt.Sprint(l["closed"])},
			})
		},
	}
}

func listCreateCmd() *cobra.Command {
	var boardID string
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a list on a board",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if boardID == "" {
				return fmt.Errorf("--board required")
			}
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.PostLists(ctx, &trello.PostListsParams{
				Name:    args[0],
				IdBoard: boardID,
			})
			if err != nil {
				return err
			}
			var l map[string]any
			if err := cmdutil.Decode(resp, &l); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "created list %s (%s)\n", strOrEmpty(l["name"]), strOrEmpty(l["id"]))
			return nil
		},
	}
	cmd.Flags().StringVarP(&boardID, "board", "b", "", "Board ID (required)")
	return cmd
}

func listArchiveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "archive <list-id>",
		Short: "Archive (close) a list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			closed := true
			resp, err := cx.Client.PutListsId(ctx, args[0], &trello.PutListsIdParams{Closed: &closed})
			if err != nil {
				return err
			}
			if err := cmdutil.Decode(resp, nil); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "archived list %s\n", args[0])
			return nil
		},
	}
}
