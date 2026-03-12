package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"

	"github.com/msc-privacy-grid-mpc-zkp/cloud-aggregator/internal/api"
	"github.com/msc-privacy-grid-mpc-zkp/cloud-aggregator/internal/zkp"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
		Name string `yaml:"name" env:"SERVER_NAME" env-default:"Server A"`
	} `yaml:"server"`

	ZKP struct {
		KeyPath  string `yaml:"key_path" env:"ZKP_KEY_PATH" env-default:"keys/verifying.key"`
		MaxLimit uint64 `yaml:"max_limit" env:"ZKP_MAX_LIMIT" env-default:"10000"`
	} `yaml:"zkp"`

	Aggregator struct {
		ExpectedMeters int `yaml:"expected_meters" env:"EXPECTED_METERS" env-default:"10"`
	} `yaml:"aggregator"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML configuration")
	flag.Parse()

	var cfg Config

	err := cleanenv.ReadConfig(*configPath, &cfg)
	if err != nil {
		log.Println("[INFO] YAML config not found, falling back to Environment variables.")
		if errEnv := cleanenv.ReadEnv(&cfg); errEnv != nil {
			log.Fatalf("[FATAL] Error loading configuration: %v", errEnv)
		}
	}

	fmt.Printf("☁️  Starting MPC Cloud Aggregator: [%s]\n", cfg.Server.Name)
	fmt.Printf("---------------------------------------------------------\n")

	verifyingKey, err := zkp.LoadVerifyingKey(cfg.ZKP.KeyPath)
	if err != nil {
		log.Fatalf("[FATAL] Failed to load verifying key: %v", err)
	}
	fmt.Println("[SECURITY] ZKP Verifying Key loaded successfully!")

	store := api.NewMemoryStore(cfg.Aggregator.ExpectedMeters)
	address := ":" + cfg.Server.Port
	api.StartServer(address, verifyingKey, store, cfg.ZKP.MaxLimit)
}
