package commands

import (
	"github.com/mrfansi/trello-cli/internal/commands/auto"
	"github.com/spf13/cobra"
)

// meCmd is a thin alias for `members get-members-id me`.
func meCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "me",
		Short: "Show the authenticated member (alias for `members get-members-id me`)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return auto.RunOp(cmd, "GET", "/members/{id}", []string{"id=me"}, nil, "")
		},
	}
}
