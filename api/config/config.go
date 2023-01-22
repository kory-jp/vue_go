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
	RedisHost string
	RedisPort string
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
		err1 := godotenv.Load("env/development.env")
		if err1 != nil {
			currentDir, _ := os.Getwd()
			if currentDir == "/app/api/infrastructure/store" {
				err2 := godotenv.Load("/app/api/env/development.env")
				if err2 != nil {
					log.Println(err1)
					log.Println(err2)
					log.Panicln(err1)
				}
			}
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
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
