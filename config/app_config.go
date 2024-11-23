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
	instance *App
	once     sync.Once
)

type App struct {
	ConfigPath string
	SecretKey  string
	Port       string
	Dsn        string
}

type dbConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" env:"DB_NAME"`
	Dsn      string
}

func NewApp(log *logger.Logger) *App {
	once.Do(func() {
		var app App

		app.ConfigPath = "./config"
		err := godotenv.Load(app.ConfigPath + "/.env")
		if err != nil {
			log.Fatalf(errors.New("can't load .env file :( "))
		}

		app.SecretKey = os.Getenv("SECRET_KEY")
		app.Port = os.Getenv("PORT")

		instance = &app
	})

	return instance

}

func Instance() *App {
	return instance
}

func getDsn(log *logger.Logger, configPath string) string {
	if os.Getenv("ENV") == "docker" {
		return os.Getenv("COMPOSE_DSN")
	}

	return getDsnFromYaml(log,configPath)

}

func getDsnFromYaml(log *logger.Logger, configPath string) string {
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/db.yaml", configPath))
	if err != nil {
		log.Fatalf(fmt.Errorf("can't load db.yaml file"))
	}

	dbConfig := &dbConfig{}

	err = yaml.Unmarshal(yamlFile, dbConfig)
	if err != nil {
		log.Fatalf(fmt.Errorf("can't unmarshal db.yaml file"))
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)
}
