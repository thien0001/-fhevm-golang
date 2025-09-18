package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RpcUrl     string
	PrivateKey string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded, reading environment variables")
	}
	return &Config{
		RpcUrl:     os.Getenv("RPC_URL"),
		PrivateKey: os.Getenv("PRIVATE_KEY"),
	}
}
