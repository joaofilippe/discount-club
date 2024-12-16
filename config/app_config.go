package appconfig

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"

	"github.com/joaofilippe/discount-club/common/logger"
)

var (
	app  *App
	once sync.Once
)

// App holds the application configuration.
type App struct {
	Environment string
	Logger      *logger.Logger
	ConfigPath  string
	SecretKey   string
	Port        string
	Dsn         string
}

type dbConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" env:"DB_NAME"`
	Dsn      string
}

// Instance returns the singleton instance of the App configuration.
func Instance(logger *logger.Logger) *App {
	if app == nil {
		return newInstance(logger)
	}

	return app
}

func newInstance(logger *logger.Logger) *App {
	once.Do(func() {
		app = &App{}
		app.Environment = os.Getenv("ENV")
		app.Logger = logger
		app.ConfigPath = "./config"

		err := godotenv.Load(app.ConfigPath + "/.env")
		if err != nil {
			app.Logger.Fatalf(errors.New("can't load .env file :( "))
		}

		app.SecretKey = os.Getenv("SECRET_KEY")
		app.Port = os.Getenv("PORT")

		setDsn()
	})

	return app
}

func setDsn() {
	if os.Getenv("ENV") == "docker" {
		app.Dsn = os.Getenv("COMPOSE_DSN")
	}

	app.Dsn = setDsnFromYaml()

}

func setDsnFromYaml() string {
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/db.yaml", app.ConfigPath))
	if err != nil {
		app.Logger.Fatalf(fmt.Errorf("can't load db.yaml file"))
	}

	dbConfig := &dbConfig{}

	err = yaml.Unmarshal(yamlFile, dbConfig)
	if err != nil {
		app.Logger.Fatalf(fmt.Errorf("can't unmarshal db.yaml file"))
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)
}
