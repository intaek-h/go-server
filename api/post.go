package api

import (
	"encoding/json"
	"net/http"

	"github.com/intaek-h/go-server/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		println("New exhibition added: ", exhibition.Title)
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// create a shell command that sends a post request to the server
// curl -X POST -d '{"title":"New Exhibition","description":"This is a new exhibition","image":"new-exhibition.png"}' -H "Content-Type: application/json" http://localhost:3333/api/exhibitions
