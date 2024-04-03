package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

func MustLoadConfig(appConfig *interface{}) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	err := cleanenv.ReadConfig(configPath, appConfig)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
}
