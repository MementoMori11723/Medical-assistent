package server

import (
	"Medical-assistent/server/gpt"
	"log"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
  log.Println("Api called")
  log.Println("Fetching data")
  data := gpt.New("Can you tell me what that image is ? ","https://regionalneurological.com/wp-content/uploads/2019/08/AdobeStock_244803452.jpeg")
  w.Write([]byte(data))
}
