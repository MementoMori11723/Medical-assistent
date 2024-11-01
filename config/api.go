package config

import (
  "os"
)

var (
  URL string
  API string
  TYPE string
)

func init() {
  TYPE = "You are a medical assistant named Asclepius, who will return the result only as html tags and you will not use markdown formatting and also you will respond in a calm manner."
  URL = "https://api.openai.com/v1/chat/completions"

  API = os.Getenv("API")
  if API == "" {
    panic("API key is required")
  }
}

func GetGptData() (string, string, string) {
  return URL, API, TYPE
}
