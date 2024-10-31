package client

import (
	"html/template"
	"log"
	"net/http"
)

var (
  pagesDir = "client/pages/"
)

func home(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFiles(
    pagesDir + "index.html",
  )

  if err != nil {
    log.Fatal(err)
  }

  if err := tmpl.Execute(w, nil); err != nil {
    log.Fatal(err)
  }
}
