package cmd

import (
	"github.com/graytonio/chamber/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [project] (args)",
	Short: "create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		// lib.CloneRepo("graytonio/anteroom-templates", "solid-ts", "")
		// lib.CloneRepo("graytonio/anteroom-templates", "", "")
		lib.CloneRepo("https://github.com/graytonio/anteroom-templates", "solid-ts", ".")
		// lib.CloneRepo("git@github.com:graytonio/anteroom-templates", "solid-ts", "")
		// lib.CloneRepo("github:graytonio/anteroom-templates", "solid-ts", "")
		return nil
	},
}
