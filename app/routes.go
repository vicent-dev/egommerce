package app

import (
	"encoding/json"
	"net/http"
)

func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

//middlewares

func (s *server) routes() {
	s.router.Use(loggingMiddleware)
	s.router.Use(jsonMiddleware)

	//ping example
	s.router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"ping\": \"pong\"}"))
	}).Methods("GET")

	s.router.HandleFunc("/user", s.regiserUser()).Methods("GET", "POST")
}
