package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func InitServer() error {
	r := mux.NewRouter()
	r.Use(jsonMiddleware)

	return http.ListenAndServe(":8080", r)
}
