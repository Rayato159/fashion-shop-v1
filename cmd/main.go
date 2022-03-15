package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rayato159/fashion-shop/v1/configs"
	"github.com/rayato159/fashion-shop/v1/internals/server"
	"github.com/rayato159/fashion-shop/v1/pkg/database"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		log.Println("error, can't load dotenv file.")
	}
	cfg := &configs.Config{}
	cfg.Database.Url = os.Getenv("DATABASE_URL")
	log.Println("dotenv has loaded successfully!")

	// Connect to database
	db, err := database.DBConnect(cfg)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	s := server.NewServer(db)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
