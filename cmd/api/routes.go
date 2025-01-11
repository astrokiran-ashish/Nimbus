package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(app.commonErrors.NotFound)
	mux.MethodNotAllowed(app.commonErrors.MethodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)

	mux.Get("/status", app.status)

	mux.Mount("/api/v1/auth", app.auth.AuthRoutes())

	return mux
}
