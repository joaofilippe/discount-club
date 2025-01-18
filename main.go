package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/joaofilippe/discount-club/app/application"
	"github.com/joaofilippe/discount-club/app/application/repositories"
	"github.com/joaofilippe/discount-club/app/application/services"
	"github.com/joaofilippe/discount-club/app/common/logger"
	"github.com/joaofilippe/discount-club/app/infra/database"
	"github.com/joaofilippe/discount-club/app/infra/webserver"
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

	discountService, _ := 
		services.NewDiscountService(repositories.NewDiscountRepo(conn))
	
		application := application.New(discountService)

	server := webserver.New(application)

	if err := server.Start(os.Getenv("PORT")); err != nil {
		logger.Fatalf(err)
	}
}
