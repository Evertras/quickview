package cmd

import (
	"github.com/Evertras/quickview/pkg/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quickview filename",
	Short: "Hot reloading viewer to see things in the browser while editing them elsewhere",
	Args:  cobra.ExactArgs(1),

	SilenceErrors: true,
	SilenceUsage:  true,

	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		s := server.New("localhost:8083", filename)

		return s.ListenAndServe()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
