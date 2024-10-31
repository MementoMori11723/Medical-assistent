package main

import (
	"Medical-assistent/client"
	"Medical-assistent/config"
	"Medical-assistent/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
  go func() {
    PORT := config.GetPort()
    client := client.New()
    server := server.New()

    http.Handle("/", client)
    http.Handle("/api", server)
    log.Println("Server running on http://localhost:"+PORT)
    http.ListenAndServe(":"+PORT, nil)
  }()
  fmt.Scanln()
}
