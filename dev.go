package main

import (
	"github.com/spf13/cobra"
)

var Dev = &cobra.Command{
	Use:   "dev",
	Short: "Development endpoint, set up all the application and run the command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx.Logger.Info("Debug entry, run whatever you want")
		return nil
	},
}

func init() {
	root.AddCommand(Dev)
}
