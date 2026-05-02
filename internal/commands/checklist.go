package commands

import (
	"fmt"

	"github.com/mrfansi/trello-cli/internal/cmdutil"
	"github.com/mrfansi/trello-cli/internal/output"
	"github.com/mrfansi/trello-cli/internal/trello"
	"github.com/spf13/cobra"
)

func checklistCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checklist",
		Short: "Manage checklists",
	}
	cmd.AddCommand(checklistLsCmd(), checklistGetCmd(), checklistCreateCmd())
	return cmd
}

func checklistLsCmd() *cobra.Command {
	var cardID string
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List checklists on a card",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if cardID == "" {
				return fmt.Errorf("--card required")
			}
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetCardsIdChecklists(ctx, cardID, &trello.GetCardsIdChecklistsParams{})
			if err != nil {
				return err
			}
			var cls []map[string]any
			if err := cmdutil.Decode(resp, &cls); err != nil {
				return err
			}
			rows := make([][]string, 0, len(cls))
			for _, cl := range cls {
				rows = append(rows, []string{
					strOrEmpty(cl["id"]),
					output.Truncate(strOrEmpty(cl["name"]), 40),
					fmt.Sprint(itemCount(cl["checkItems"])),
				})
			}
			return cx.Renderer.Table([]string{"ID", "Name", "Items"}, rows)
		},
	}
	cmd.Flags().StringVarP(&cardID, "card", "c", "", "Card ID (required)")
	return cmd
}

func checklistGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <checklist-id>",
		Short: "Get a checklist by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetChecklistsId(ctx, args[0], &trello.GetChecklistsIdParams{})
			if err != nil {
				return err
			}
			var cl map[string]any
			if err := cmdutil.Decode(resp, &cl); err != nil {
				return err
			}
			if jsonOutput {
				return cx.Renderer.Raw(cl)
			}
			return cx.Renderer.KeyValue([][2]string{
				{"ID", strOrEmpty(cl["id"])},
				{"Name", strOrEmpty(cl["name"])},
				{"Card ID", strOrEmpty(cl["idCard"])},
				{"Items", fmt.Sprint(itemCount(cl["checkItems"]))},
			})
		},
	}
}

func checklistCreateCmd() *cobra.Command {
	var cardID string
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a checklist on a card",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if cardID == "" {
				return fmt.Errorf("--card required")
			}
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			name := args[0]
			resp, err := cx.Client.PostChecklists(ctx, &trello.PostChecklistsParams{
				IdCard: cardID,
				Name:   &name,
			})
			if err != nil {
				return err
			}
			var cl map[string]any
			if err := cmdutil.Decode(resp, &cl); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "created checklist %s (%s)\n", strOrEmpty(cl["name"]), strOrEmpty(cl["id"]))
			return nil
		},
	}
	cmd.Flags().StringVarP(&cardID, "card", "c", "", "Card ID (required)")
	return cmd
}

func itemCount(v any) int {
	if v == nil {
		return 0
	}
	if arr, ok := v.([]any); ok {
		return len(arr)
	}
	return 0
}
