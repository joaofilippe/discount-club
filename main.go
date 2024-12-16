package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/joaofilippe/discount-club/commands"
	"github.com/joaofilippe/discount-club/common/logger"
	"github.com/joaofilippe/discount-club/config"
	"github.com/joaofilippe/discount-club/infra/database"
	"github.com/joaofilippe/discount-club/infra/database/migrations"
)

func main() {
	logger := logger.New()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(commands.CreateEnvCommand())

	if err := rootCmd.Execute(); err != nil {
		logger.Fatalf(err)
	}

	app := appconfig.Instance(logger)

	conn := database.New(logger, app)
	if err := migrations.Up(conn); err != nil {
		_ = migrations.Down(conn)
		logger.Fatalf(fmt.Errorf("can't run migrations: %w", err))
	}
}
