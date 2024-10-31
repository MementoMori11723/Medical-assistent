package server

import "net/http"

func api(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("API"))
}
