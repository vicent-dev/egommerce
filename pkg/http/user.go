package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vicent-dev/egommerce/pkg/user/register"
)

var (
	registerService register.Service
)

func handleUser(r *mux.Router, c context.Context) {
	registerService = c.Value("service_register_user").(register.Service)
	r.HandleFunc("/user/register", registerUser).Methods("POST")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var u register.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		errorResponse(w, err)
		return
	}

	if err := registerService.Register(u); err != nil {
		errorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return

}
