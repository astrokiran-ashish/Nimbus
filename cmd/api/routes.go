package main

import (
	"net/http"

	commonErrors "github.com/astrokiran/nimbus/internal/common/errors"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(commonErrors.NotFound)
	mux.MethodNotAllowed(commonErrors.MethodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)

	mux.Get("/status", app.status)

	mux.Mount("/api/v1/auth", app.auth.AuthRoutes())
	mux.Mount("/api/v1/consultant", app.consultant.ConsultantRoutes())
	mux.Mount("/api/v1/consultation", app.consultation.ConsultationRoutes())

	return mux
}
