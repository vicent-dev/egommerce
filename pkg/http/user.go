package http

import (
	"context"
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

	if err := register.Service.Register(); err != nil {
		return
	}
	return
}
