package main

import (
	"net/http"

	container "github.com/gabrielbsx/wyd-go/internal/di"
	"github.com/gabrielbsx/wyd-go/presenter"

	"github.com/go-chi/chi"
)

func main() {
	app := container.NewAppContainer()

	r := chi.NewRouter()

	r.Route("/api/v1/users", func(r chi.Router) {
		presenter.NewAppContainerWrapper(app, r)
	})

	app.Logger.Infof("Server started at :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		app.Logger.Errorf("error: %v", err)
	}
}
