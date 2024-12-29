package main

import (
	"github.com/spf13/cobra"

	"github.com/joaofilippe/discount-club/commands"
	"github.com/joaofilippe/discount-club/app/common/logger"
	appconfig "github.com/joaofilippe/discount-club/config"
	"github.com/joaofilippe/discount-club/app/infra/database"
)

func main() {
	logger := logger.New()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(commands.RunAppCommand())

	if err := rootCmd.Execute(); err != nil {
		logger.Fatalf(err)
	}

	app := appconfig.Instance(logger)

	conn := database.New(logger, app)
	if err := conn.Get().Ping(); err != nil {
		logger.Fatalf(err)
	}

	conn.RunMigrations()

}
