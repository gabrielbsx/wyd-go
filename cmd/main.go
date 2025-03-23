package main

import (
	"net/http"

	"github.com/gabrielbsx/wyd-go/infra/logger"
	"github.com/gabrielbsx/wyd-go/presenter"

	"github.com/go-chi/chi"
)

func main() {
	logger := logger.NewLogger()

	r := chi.NewRouter()

	r.Route("/api/v1/users", func(r chi.Router) {
		presenter.SignUpRoute(r)
	})

	logger.Infof("Server started at :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Errorf("error: %v", err)
	}
}
