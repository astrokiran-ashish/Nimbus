package auth

import "github.com/go-chi/chi/v5"

func (auth *Auth) AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", auth.LoginViaOTP)

	return r

}
