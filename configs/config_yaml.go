package configs

import (
	"github.com/Ptsecuritydit/libs/redisdb"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

var ServiceConfig AppConfig

func init() {
	MustLoadConfig()
}

type AppConfig struct {
	RedisConfig    redisdb.Conf      `yaml:"redis"`
	KafkaMq        KafkaConfig       `yaml:"kafka"`
	UseRedisEvents bool              `yaml:"useRedisEvents"`
	Namespaces     map[string]string `yaml:"namespaces"`
	MongoConfig    `yaml:"mongodb"`
	UseKafka       bool `yaml:"useKafka"`
	HttpServer     `yaml:"http"`
}

type KafkaConfig struct {
	Config             []map[string]string `yaml:"config"`
	AdditionalSettings string              `yaml:"additionalSettings"`
	Topics             []string            `yaml:"topics"`
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type MongoConfig struct {
	ConnectionString string `yaml:"connectionString" env-default:"mongodb://localhost:27017"`
	Database         string `yaml:"database" env-default:"test"`
	Collection       string `yaml:"collection" env-default:"test"`
}

type RabbitConfig struct {
	ConnectionString string `yaml:"connectionString" env-default:"amqp://localhost:5672/"`
}

func MustLoadConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	err := cleanenv.ReadConfig(configPath, &ServiceConfig)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
}
