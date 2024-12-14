package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// GetEnvFlag returns a cobra command for environment flag
func GetEnvFlag() *cobra.Command {
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

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(GetEnvFlag())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	fmt.Println("enviroment: ", os.Getenv("ENV"))
}
