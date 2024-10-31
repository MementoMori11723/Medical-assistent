package config

import (
  "os"
  "github.com/joho/godotenv"
)

var (
  URL string
  API string
  TYPE string
)

func init() {
  TYPE = "You are a medical assistant named Asclepius, which will return the result as pure plain text and no formatting so that it can be easily read by the user."
  URL = "https://api.openai.com/v1/chat/completions"
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  API = os.Getenv("API")
  if API == "" {
    panic("API key is required")
  }
}

func GetGptData() (string, string, string) {
  return URL, API, TYPE
}
