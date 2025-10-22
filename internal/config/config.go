package config

import (
	"flag"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
    Addr string `yaml:"addr" env:"HTTP_ADDR" env-required:"true" env-default:":8080"`
}

type Config struct {
    Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"` 
    StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"` 
    HTTPServer  HTTPServer `yaml:"http_server"`
}


func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "" , "Path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not provided")
		}
	}

	if _, err := os.Stat((configPath)); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg 
	
}