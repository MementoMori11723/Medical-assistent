package config

import (
	"log"
	"os"

)

var (
  PORT string
)

func init() {
  PORT = os.Getenv("PORT")
  if PORT == "" {
    log.Println("PORT not found in .env file, using default value")
    PORT = "8080"
  }
}

func GetPort() string {
  return PORT
}
