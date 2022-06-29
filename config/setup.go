package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	MONGO_HOST     string `mapstructure:"MONGO_HOST"`
	MONGO_PORT     string `mapstructure:"MONGO_PORT"`
	MONGO_USER     string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	MONGO_PASS     string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	MONGO_DATABASE string `mapstructure:"MONGO_DATABASE"`
	MODE           string `mapstructure:"MODE"`
}

func getConfig() *Config {
	log.Println("Loading config...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var conf *Config
	envMap := make(map[string]interface{})
	for _, env := range os.Environ() {
		key := env[:strings.Index(env, "=")]
		value := env[strings.Index(env, "=")+1:]
		envMap[key] = value
	}

	err = mapstructure.Decode(envMap, &conf)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return conf
}

var Cfg *Config

func Setup() {

	Cfg = getConfig()

	log.Println("Config loaded")

	newMongoClient()
}
