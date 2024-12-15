package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/joaofilippe/discount-club/common/logger"
	appconfig "github.com/joaofilippe/discount-club/config"
	"github.com/joaofilippe/discount-club/infra/database"
	"github.com/joaofilippe/discount-club/infra/database/migrations"
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

	logger := logger.New()
	app := appconfig.Instance(logger)

	conn := database.New(logger, app)

	if err := conn.Get().Ping(); err != nil {
		panic(fmt.Errorf("can't ping database: %w", err))

	}

	if err := migrations.Up(conn); err != nil {
		migrations.Down(conn)
		panic(fmt.Errorf("can't run migrations: %w", err))
	}
}
