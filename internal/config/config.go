package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

func MustLoad() *Config {
	// ошибка, не читается энв
	// configPath := os.Getenv("CONFIG_PATH")
	// if configPath == "" {
	// 	log.Fatal("CONFIG_PATH is not set")
	// }

	configPath := "../../config/local.yaml"

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file does not exist: %v", err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config file - %v", err)
	}
	return &cfg
}
