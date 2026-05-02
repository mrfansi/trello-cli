package commands

import "github.com/spf13/cobra"

var jsonOutput bool

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:           "trello-cli",
		Short:         "Trello CLI",
		Long:          "Curated CLI over the Trello REST API.\n\nAuth: TRELLO_API_KEY + TRELLO_TOKEN env vars, or ~/.trello-cli/config.yaml.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output raw JSON")

	root.AddCommand(
		boardCmd(),
		listCmd(),
		cardCmd(),
		checklistCmd(),
		memberCmd(),
		meCmd(),
	)
	return root
}
