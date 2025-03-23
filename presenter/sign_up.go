package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielbsx/wyd-go/app/handlers/identity"
	container "github.com/gabrielbsx/wyd-go/internal/di"
	"github.com/go-chi/chi"
)

func SignUpRoute(r chi.Router) {
	app := container.NewAppContainer()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		body := identity.SignUpUserCommand{}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			sendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = app.SignUp.Handle(body)

		if err != nil {
			sendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		sendSuccess(w, nil, http.StatusCreated)
	})
}
