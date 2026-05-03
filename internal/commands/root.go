package commands

import (
	"github.com/mrfansi/trecli/internal/commands/auto"
	"github.com/mrfansi/trecli/internal/version"
	"github.com/spf13/cobra"
)

var jsonOutput bool

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:           "trecli",
		Short:         "Trello CLI",
		Long:          "CLI over the Trello REST API. Resource groups are auto-generated from openapi.json. Use `raw` for ad-hoc requests.\n\nAuth: TRELLO_API_KEY + TRELLO_TOKEN env vars, or ~/.trecli/config.yaml.",
		Version:       version.String(),
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.SetVersionTemplate("trecli {{.Version}}\n")
	root.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output raw JSON (no-op for raw/auto commands which always emit JSON)")

	root.AddCommand(meCmd())
	root.AddCommand(rawCmd())
	root.AddCommand(auto.Groups()...)
	return root
}
