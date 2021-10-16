package app

import (
	"encoding/json"
	"net/http"

	"github.com/en-vee/alog"
)

func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

//middlewares
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		alog.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *server) routes() {
	s.router.Use(loggingMiddleware)
	s.router.Use(jsonMiddleware)

	//ping example
	s.router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"ping\": \"pong\"}"))
	}).Methods("GET")

	s.router.HandleFunc("/user", s.regiserUser()).Methods("GET", "POST")
}
