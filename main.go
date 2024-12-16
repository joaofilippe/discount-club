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
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(commands.CreateEnvCommand())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	logger := logger.New()
	app := appconfig.Instance(logger)

	conn := database.New(logger, app)

	if err := conn.Get().Ping(); err != nil {
		logger.Errorf(fmt.Errorf("can't ping database: %w", err))
	}

	if err := migrations.Up(conn); err != nil {
		_ = migrations.Down(conn)
		logger.Fatalf(fmt.Errorf("can't run migrations: %w", err))
	}
}
