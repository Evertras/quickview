package cmd

import (
	"github.com/Evertras/quickview/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "quickview filename",
	Short: "Hot reloading viewer to see things in the browser while editing them elsewhere",
	Args:  cobra.ExactArgs(1),

	SilenceErrors: true,
	SilenceUsage:  true,

	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		s := server.New(config.Address, filename)

		return s.ListenAndServe()
	},
}

func init() {
	rootCmd.Flags().StringP(configKeyAddress, "a", "localhost:8386", "The address to host on.")

	err := viper.BindPFlags(rootCmd.Flags())

	if err != nil {
		panic(err)
	}
}
