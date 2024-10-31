package server

import (
	"Medical-assistent/server/gpt"
	"encoding/json"
	"io"
	"net/http"
)

type Body struct {
	Text string `json:"text,omitempty"`
	Url  string `json:"url,omitempty"`
}

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
