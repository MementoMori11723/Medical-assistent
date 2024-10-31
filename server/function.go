package server

import (
	"Medical-assistent/server/gpt"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
  data := gpt.New("What should i do if i have a fever?","")
  w.Write([]byte(data))
}
