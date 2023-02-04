package config

import (
	"log"
	"os"

	"github.com/kory-jp/vue_go/api/utiles"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	Env       string
	Port      string
	LogFile   string
	SQLDriver string
	UserName  string
	Password  string
	DBHost    string
	DBPort    string
	DBName    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utiles.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		err := godotenv.Load("production.env")
		if err != nil {
			log.Println(err)
			log.Panicln(err)
		}
	} else if env == "development" {
		err := godotenv.Load("env/development.env")
		if err != nil {
			log.Println(err)
			log.Panicln(err)
		}
	}

	Config = ConfigList{
		Env:       os.Getenv("GO_ENV"),
		Port:      os.Getenv("API_PORT"),
		LogFile:   os.Getenv("LOG_FILE"),
		SQLDriver: os.Getenv("DRIVER"),
		UserName:  os.Getenv("USER_NAME"),
		Password:  os.Getenv("PASSWORD"),
		DBHost:    os.Getenv("HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
	}
}
