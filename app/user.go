package app

import (
	"net/http"

	"github.com/vicent-dev/egommerce/user"
)

func (s *server) regiserUser() http.HandlerFunc {
	//first time handler called (preload x)
	return func(w http.ResponseWriter, r *http.Request) {
		//call use case from user context
		err := user.RegisterFakeUser()
		if err != nil {
			errorResponse(w, err)
			return
		}

		w.Write([]byte("{\"success\": true}"))
	}
}
