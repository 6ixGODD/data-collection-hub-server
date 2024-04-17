package cmd

import (
	"github.com/spf13/cobra"
)

// TODO: If configuration params are given, use them overwriting the default values and environment variables

var RootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "cmd is a tool to generate Cobra-based application",
	Long: `cmd is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra-based application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

	},
}

func Execute() error {
	return RootCmd.Execute()
}
