package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/en-vee/alog"
	"github.com/gorilla/mux"
)

func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

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

func InitServer(ctx context.Context) error {
	r := mux.NewRouter()
	r.Use(jsonMiddleware)
	r.Use(loggingMiddleware)

	handleUser(r, ctx)

	return http.ListenAndServe(":8080", r)
}
