package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
  PORT string
)

func init() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  PORT = os.Getenv("PORT")
  if PORT == "" {
    log.Println("PORT not found in .env file, using default value")
    PORT = "8080"
  }
}

func GetPort() string {
  return PORT
}
