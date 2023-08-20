package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quickview",
	Short: "Hot reloading viewer to see things in the browser while editing them elsewhere",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi I ran")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
