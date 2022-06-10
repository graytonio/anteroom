package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [project] (args)",
	Short: "create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
