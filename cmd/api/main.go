package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thien0001/fhevm-go-demo/pkg/config"
	"github.com/thien0001/fhevm-go-demo/pkg/fhevm"
	"github.com/thien0001/fhevm-go-demo/pkg/handlers"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file loaded, reading environment variables")
	}
}

func main() {
	rpcURL := os.Getenv("RPC_URL")
	privateKey := os.Getenv("PRIVATE_KEY")
	cfg := config.Load()
	if cfg.RpcUrl == "" || cfg.PrivateKey == "" {
		log.Println("Please set RPC_URL and PRIVATE_KEY in env")
		os.Exit(1)
	}
	fmt.Println("RPC_URL:", rpcURL)
	fmt.Println("PRIVATE_KEY:", privateKey[:10]+"...")
	client, err := fhevm.NewClient(cfg.RpcUrl, cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Static("/", "./static")
	api := handlers.NewApiHandler(client)
	app.Post("/api/encrypt", api.Encrypt)
	app.Post("/api/send", api.Send)
	app.Get("/api/result/:txHash", api.Result)

	log.Println("Server listening on :3000")
	log.Fatal(app.Listen(":3000"))
}
