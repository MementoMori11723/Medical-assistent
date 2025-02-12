package main

import (
	"Medical-assistent/config"
	"Medical-assistent/server"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	PORT := config.GetPort()
	handler := server.New()
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: handler,
	}
	slog.Info("Server running on http://localhost:" + PORT)
  err := server.ListenAndServe()
  if err != nil {
    slog.Error(err.Error())
    os.Exit(1)
  }
}
