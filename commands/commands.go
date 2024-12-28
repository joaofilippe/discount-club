package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RunAppCommand() *cobra.Command {
	var env string
	var resetDB bool

	cmd := &cobra.Command{
		Use:   "app",
		Short: "Run application with arguments",
		Run: func(cmd *cobra.Command, args []string) {
			os.Setenv("ENV", env)
			os.Setenv("RESET_DB", fmt.Sprint(resetDB))
		},
	}

	cmd.Flags().StringVarP(&env, "env", "e", "dev", "Environment flag")
	cmd.Flags().BoolVarP(&resetDB, "reset-db", "r", false, "Reset Database")
	return cmd
}
