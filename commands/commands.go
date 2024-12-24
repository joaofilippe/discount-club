package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func CreateEnvCommand() *cobra.Command {
	var env string
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Get environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			os.Setenv("ENV", env)
		},
	}

	cmd.Flags().StringVarP(&env, "env", "e", "dev", "Environment flag")

	return cmd
}
