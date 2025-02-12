package server

import (
	"Medical-assistent/server/gpt"
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Body struct {
	Text string `json:"text,omitempty"`
	Url  string `json:"url,omitempty"`
}

//go:embed pages/*
var pages embed.FS

func api(w http.ResponseWriter, r *http.Request) {
	var body Body
	ioBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(ioBody, &body)
	data := gpt.New(body.Text, body.Url)
	w.Write([]byte(data))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	temp, err := template.ParseFS(pages, "pages/index.html")
	if err != nil {
		http.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}
	if err := temp.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
