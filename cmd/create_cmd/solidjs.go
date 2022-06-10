package create_cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var SolidJSCmd = &cobra.Command{
	Use:   "solidjs [path]",
	Short: "generate a solidjs project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a path argument")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmds := []string{fmt.Sprintf("npx degit solidjs/templates/ts %s")}
	},
}
