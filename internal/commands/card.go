package commands

import (
	"fmt"

	"github.com/mrfansi/trello-cli/internal/cmdutil"
	"github.com/mrfansi/trello-cli/internal/output"
	"github.com/mrfansi/trello-cli/internal/trello"
	"github.com/spf13/cobra"
)

func cardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Manage cards",
	}
	cmd.AddCommand(cardLsCmd(), cardGetCmd(), cardCreateCmd(), cardUpdateCmd(), cardDeleteCmd())
	return cmd
}

func cardLsCmd() *cobra.Command {
	var (
		listID  string
		boardID string
	)
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List cards on a board or list",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if (listID == "") == (boardID == "") {
				return fmt.Errorf("exactly one of --list or --board required")
			}
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			var (
				cards []map[string]any
				err2  error
			)
			if listID != "" {
				resp, e := cx.Client.GetListsIdCards(ctx, listID)
				if e != nil {
					return e
				}
				err2 = cmdutil.Decode(resp, &cards)
			} else {
				resp, e := cx.Client.GetBoardsIdCards(ctx, boardID)
				if e != nil {
					return e
				}
				err2 = cmdutil.Decode(resp, &cards)
			}
			if err2 != nil {
				return err2
			}
			rows := make([][]string, 0, len(cards))
			for _, c := range cards {
				rows = append(rows, []string{
					strOrEmpty(c["id"]),
					output.Truncate(strOrEmpty(c["name"]), 50),
					strOrEmpty(c["idList"]),
					strOrEmpty(c["shortUrl"]),
				})
			}
			return cx.Renderer.Table([]string{"ID", "Name", "List ID", "URL"}, rows)
		},
	}
	cmd.Flags().StringVarP(&listID, "list", "l", "", "List ID")
	cmd.Flags().StringVarP(&boardID, "board", "b", "", "Board ID")
	return cmd
}

func cardGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <card-id>",
		Short: "Get a card by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.GetCardsId(ctx, args[0], &trello.GetCardsIdParams{})
			if err != nil {
				return err
			}
			var c map[string]any
			if err := cmdutil.Decode(resp, &c); err != nil {
				return err
			}
			if jsonOutput {
				return cx.Renderer.Raw(c)
			}
			return cx.Renderer.KeyValue([][2]string{
				{"ID", strOrEmpty(c["id"])},
				{"Name", strOrEmpty(c["name"])},
				{"Description", strOrEmpty(c["desc"])},
				{"Due", strOrEmpty(c["due"])},
				{"Closed", fmt.Sprint(c["closed"])},
				{"List ID", strOrEmpty(c["idList"])},
				{"Board ID", strOrEmpty(c["idBoard"])},
				{"URL", strOrEmpty(c["shortUrl"])},
			})
		},
	}
}

func cardCreateCmd() *cobra.Command {
	var (
		listID string
		desc   string
	)
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a card",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if listID == "" {
				return fmt.Errorf("--list required")
			}
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			name := args[0]
			params := &trello.PostCardsParams{
				Name:   &name,
				IdList: listID,
			}
			if desc != "" {
				params.Desc = &desc
			}
			resp, err := cx.Client.PostCards(ctx, params)
			if err != nil {
				return err
			}
			var c map[string]any
			if err := cmdutil.Decode(resp, &c); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "created card %s (%s)\n", strOrEmpty(c["name"]), strOrEmpty(c["id"]))
			return nil
		},
	}
	cmd.Flags().StringVarP(&listID, "list", "l", "", "List ID (required)")
	cmd.Flags().StringVar(&desc, "desc", "", "Card description")
	return cmd
}

func cardUpdateCmd() *cobra.Command {
	var (
		name   string
		desc   string
		listID string
		closed bool
		closeF bool
	)
	cmd := &cobra.Command{
		Use:   "update <card-id>",
		Short: "Update a card",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			params := &trello.PutCardsIdParams{}
			if name != "" {
				params.Name = &name
			}
			if desc != "" {
				params.Desc = &desc
			}
			if listID != "" {
				v := trello.TrelloID(listID)
				params.IdList = &v
			}
			if cmd.Flags().Changed("closed") {
				closeF = closed
				params.Closed = &closeF
			}

			resp, err := cx.Client.PutCardsId(ctx, args[0], params)
			if err != nil {
				return err
			}
			if err := cmdutil.Decode(resp, nil); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "updated card %s\n", args[0])
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "New name")
	cmd.Flags().StringVar(&desc, "desc", "", "New description")
	cmd.Flags().StringVarP(&listID, "list", "l", "", "Move to list ID")
	cmd.Flags().BoolVar(&closed, "closed", false, "Set archived state")
	return cmd
}

func cardDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rm <card-id>",
		Short: "Delete a card",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cx, err := cmdutil.Build(jsonOutput)
			if err != nil {
				return err
			}
			ctx, cancel := cmdutil.Context()
			defer cancel()

			resp, err := cx.Client.DeleteCardsId(ctx, args[0])
			if err != nil {
				return err
			}
			if err := cmdutil.Decode(resp, nil); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "deleted card %s\n", args[0])
			return nil
		},
	}
}
