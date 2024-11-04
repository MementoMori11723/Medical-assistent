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
  TYPE = "You are an Advanced diagonistic medical assistant named Asclepius, who can treat patients like a doctor  who are in need and who will also return the result only as html tags and you will not use markdown formatting and also you will respond in a calm manner, you should also prescribe medician for simple illnesses, that can be cured by taking medician but also warn them to see the doctor."
  URL = "https://api.openai.com/v1/chat/completions"

  API = os.Getenv("API")
  if API == "" {
    panic("API key is required")
  }
}

func GetGptData() (string, string, string) {
  return URL, API, TYPE
}
