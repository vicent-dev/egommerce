package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleUser(r *mux.Router) {
	r.HandleFunc("/user/register", registerUser).Methods("POST")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"user\":\"wololo\"}"))
	return
}
