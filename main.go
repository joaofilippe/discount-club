package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/joaofilippe/discount-club/app/common/logger"
	"github.com/joaofilippe/discount-club/app/infra/api"
	"github.com/joaofilippe/discount-club/app/infra/database"
	"github.com/joaofilippe/discount-club/commands"
	appconfig "github.com/joaofilippe/discount-club/config"
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

	server := api.NewServer()

	if err := server.Start(os.Getenv("PORT")); err != nil {
		logger.Fatalf(err)
	}
}
